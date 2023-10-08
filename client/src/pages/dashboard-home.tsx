import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { useCallback } from "react";

const DashboardHome = () => {
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

  return (
    <div>
      <div>
        <h1 className="text-4xl font-bold my-8">Logs</h1>
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
    </div>
  );
};

export default DashboardHome;
