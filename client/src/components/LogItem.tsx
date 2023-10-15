const LogItem = ({ log }: { log: any }) => {
  return (
    <details className="m-1" open>
      <summary className="flex items-center justify-between cursor-pointer">
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
  );
};

export default LogItem;
