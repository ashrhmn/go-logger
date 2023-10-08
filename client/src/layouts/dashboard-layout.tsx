import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import md5 from "md5";
import { useMemo } from "react";
import { Link, Outlet, useNavigate } from "react-router-dom";

const DashboardLayout = () => {
  const navigate = useNavigate();
  const { data: loggedInUser } = useQuery({
    queryKey: ["whoami"],
    queryFn: () => axios.get("/api/auth/whoami").then((res) => res.data),
  });
  const avatarUrl = useMemo(
    () =>
      `https://www.gravatar.com/avatar/${md5(
        loggedInUser?.email || "ghost@email.com"
      )}?s=2048`,
    [loggedInUser?.email]
  );

  const handleLogout = () =>
    axios.delete("/api/auth").then(() => navigate("/login"));

  return (
    <>
      <div className="navbar bg-base-100">
        <div className="flex-1">
          <Link to={"/dashboard"} className="btn btn-ghost normal-case text-xl">
            Go-Logger
          </Link>
        </div>
        <div className="flex-none gap-2">
          <div className="form-control">
            <input
              type="text"
              placeholder="Search"
              className="input input-bordered w-24 md:w-auto"
            />
          </div>
          <div className="dropdown dropdown-end">
            <label tabIndex={0} className="btn btn-ghost btn-circle avatar">
              <div className="w-10 h-10 rounded-full flex justify-center items-center">
                <img src={avatarUrl} />
              </div>
            </label>
            <ul
              tabIndex={0}
              className="mt-3 z-[10] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52"
            >
              <li>
                <a className="justify-between">
                  Profile
                  <span className="badge">Soon</span>
                </a>
              </li>
              <li>
                <Link to="/dashboard/settings">Settings</Link>
              </li>
              <li>
                <a onClick={handleLogout}>Logout</a>
              </li>
            </ul>
          </div>
        </div>
      </div>
      <Outlet />
    </>
  );
};

export default DashboardLayout;
