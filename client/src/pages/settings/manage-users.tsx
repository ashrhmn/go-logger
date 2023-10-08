/* eslint-disable @typescript-eslint/no-explicit-any */

import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import md5 from "md5";
import EditUserForm from "../../components/Settings/ManageUsers/EditUserForm";

const ManageUsers = () => {
  const { data: users } = useQuery({
    queryKey: ["users"],
    queryFn: () => axios.get("/api/users").then((res) => res.data),
  });

  return (
    <div className="">
      <h1>Manage Users</h1>
      <div className="max-h-[80vh] w-full overflow-x-auto">
        <table className="table table-xs table-pin-rows table-pin-cols overflow-x-auto">
          <thead>
            <tr>
              <th></th>
              <td>Username</td>
              <td>Email</td>
              <td>Avatar</td>
              <td>Name</td>
              <td>Permissions</td>
              <td>Created By</td>
              <td>Created At</td>
              <td>Last Updated By</td>
              <td>Last Updated At</td>
              <td>Actions</td>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {users?.map((user: any, i: number) => (
              <tr key={user.id + i}>
                <th>{i + 1}</th>
                <td>{user.username}</td>
                <td>{user.email}</td>
                <td>
                  <img
                    className="w-10 h-10 rounded-full"
                    src={`https://www.gravatar.com/avatar/${md5(
                      user.email
                    )}?s=2048`}
                    alt=""
                  />
                </td>
                <td>{user.firstName + " " + user.lastName}</td>
                <td>{user.permissions.join(", ")}</td>
                <td>{user.createdBy}</td>
                <td>{new Date(user.createdAt * 1000).toLocaleString()}</td>
                <td>{user.updatedBy}</td>
                <td>{new Date(user.updatedAt * 1000).toLocaleString()}</td>
                <td className="flex items-center gap-1">
                  <button
                    onClick={() =>
                      (
                        document.getElementById(
                          `user_edit_modal_${user.id}`
                        ) as HTMLDialogElement
                      )?.showModal()
                    }
                    className="btn btn-sm btn-accent"
                  >
                    Edit
                  </button>
                  <button className="btn btn-sm btn-warning">
                    {user.deletedAt === 0 ? "Delete" : "Delete Permanently"}
                  </button>
                </td>
                <th>{i + 1}</th>
                <EditUserForm user={user} />
              </tr>
            ))}
          </tbody>
          <tfoot>
            <tr>
              <th></th>
              <td>Username</td>
              <td>Email</td>
              <td>Avatar</td>
              <td>Name</td>
              <td>Permissions</td>
              <td>Created By</td>
              <td>Created At</td>
              <td>Last Updated By</td>
              <td>Last Updated At</td>
              <td>Actions</td>
              <th></th>
            </tr>
          </tfoot>
        </table>
      </div>
    </div>
  );
};

export default ManageUsers;
