import { Link, Outlet } from "react-router-dom";

const SettingsLayout = () => {
  return (
    <main className="flex">
      <aside className="w-96 flex flex-col p-4">
        <h1 className="text-4xl font-bold my-8">Settings</h1>
        <ul className="flex flex-col">
          <li>
            <Link to="manage-users">Manage Users</Link>
          </li>
        </ul>
      </aside>
      <Outlet />
    </main>
  );
};

export default SettingsLayout;
