import { Link, Outlet } from "react-router-dom";

const SettingsLayout = () => {
  return (
    <div className="flex">
      <aside className="w-[40%] max-w-[200px] hidden flex-col p-4 xs:flex bg-base-200 min-h-[80vh]">
        <h1 className="text-4xl font-bold my-8">Settings</h1>
        <ul className="flex flex-col">
          <li>
            <Link to="manage-users">Manage Users</Link>
          </li>
        </ul>
      </aside>
      <div className="w-[100vw] xs:min-w-[60%] p-6">
        <Outlet />
      </div>
    </div>
  );
};

export default SettingsLayout;
