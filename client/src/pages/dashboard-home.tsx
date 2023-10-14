import { useQuery, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useCallback, useEffect, useMemo, useState } from "react";

const DashboardHome = () => {
  const [limit, setLimit] = useState(40);
  const queryClient = useQueryClient();
  const { data: logLevels } = useQuery({
    queryKey: ["all-log-levels"],
    queryFn: () =>
      axios
        .get("/api/logging/all-log-levels")
        .then((res) => res.data as string[]),
  });
  const { data: selectedLogLevels, refetch: refetchSelectedLogLevels } =
    useQuery({
      queryKey: ["selected-log-levels"],
      queryFn: () =>
        axios
          .get("/api/logging/selected-log-levels")
          .then((res) => res.data as string[]),
    });

  const updateSelectedLogLevels = useCallback(
    (level: string, checked: boolean) => {
      axios
        .patch(
          "/api/logging/selected-log-levels",
          checked
            ? Array.from(new Set([...(selectedLogLevels || []), level]))
            : (selectedLogLevels || []).filter((l) => l !== level)
        )
        .then(() => refetchSelectedLogLevels())
        .catch((err) => console.error(err));
    },
    [refetchSelectedLogLevels, selectedLogLevels]
  );

  const commaSeparatedSelectedLogLevels = useMemo(
    () => (!selectedLogLevels ? undefined : selectedLogLevels.sort().join(",")),
    [selectedLogLevels]
  );

  const logQueryKey = useMemo(
    () => ["logs", commaSeparatedSelectedLogLevels, limit],
    [commaSeparatedSelectedLogLevels, limit]
  );

  const { data: logs } = useQuery({
    queryKey: logQueryKey,
    queryFn: () =>
      axios
        .get("/api/logging/logs", {
          params: { levels: commaSeparatedSelectedLogLevels, limit },
        })
        .then((res) => res.data),
    keepPreviousData: true,
  });

  console.log({ logs });

  const onLogMessage = useCallback(
    (log: any) => {
      if (!selectedLogLevels?.includes(log.level)) return;
      queryClient.setQueryData(logQueryKey, (oldLogs: any) => [
        log,
        ...(oldLogs || []),
      ]);
    },
    [logQueryKey, queryClient, selectedLogLevels]
  );

  useEffect(() => {
    const ws = new WebSocket(
      `${window.location.protocol.includes("https") ? "wss" : "ws"}://${
        window.location.host
      }/api/socket/ws/logs`
    );
    ws.onopen = () => {
      console.log("connected");
    };
    ws.onmessage = (e) => {
      const log = JSON.parse(e.data);
      console.log(log);
      onLogMessage(log);
    };
    ws.onclose = () => {
      console.log("disconnected");
    };
    return () => {
      ws.close();
    };
  }, [onLogMessage]);

  return (
    <div>
      <div className="my-8">
        <h1 className="text-4xl font-bold my-4">Logs</h1>
        <h1>Filters</h1>
        <div className="flex gap-3 flex-wrap items-center">
          {logLevels?.map((level) => (
            <div
              key={level}
              className="form-control w-36 bg-base-300 rounded p-1 text-sm"
            >
              <label className="label cursor-pointer">
                <span className="label-text">{level.toUpperCase()}</span>
                <input
                  type="checkbox"
                  className="checkbox checkbox-xs"
                  checked={selectedLogLevels?.includes(level) || false}
                  onChange={(e) =>
                    updateSelectedLogLevels(level, e.target.checked)
                  }
                />
              </label>
            </div>
          ))}
        </div>
      </div>
      <div className="w-full h-[60vh] overflow-y-auto flex flex-col-reverse p-4">
        {logs?.map((log: any) => (
          <details open key={log.id}>
            <summary className="flex items-center justify-between">
              <div className="text-sm text-gray-500 flex items-center gap-3">
                <span>{new Date(log.timestamp * 1000).toLocaleString()}</span>
                {!!log.context && <span>Context : {log.context}</span>}
                {!!log.note && <span>Note : {log.note}</span>}
              </div>
              <div className="text-sm text-gray-500">{log.level}</div>
            </summary>
            <pre className="text-sm text-gray-500">
              {JSON.stringify(log.payload, null, 2)}
            </pre>
          </details>
        ))}
        <button onClick={() => setLimit((l) => l + 10)}>Load More</button>
      </div>
    </div>
  );
};

export default DashboardHome;
