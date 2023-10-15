import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import md5 from "md5";
import AddUserModal from "../../components/Settings/ManageUsers/AddUserModal";
import DeleteUserModal from "../../components/Settings/ManageUsers/DeleteUserModal";
import EditUserForm from "../../components/Settings/ManageUsers/EditUserForm";
import { clx } from "../../utils/jsx.utils";

const ManageUsers = () => {
  const { data: users, refetch: refetchUsers } = useQuery({
    queryKey: ["all-users"],
    queryFn: () => axios.get("/api/users").then((res) => res.data),
  });

  return (
    <div className="">
      <AddUserModal />
      <div className="flex justify-between items-center my-10">
        <h1 className="text-xl font-bold">Manage Users</h1>
        <button
          onClick={() =>
            (
              document.getElementById(`user_add_modal`) as HTMLDialogElement
            )?.showModal()
          }
          className="btn btn-sm btn-primary"
        >
          Add New User
        </button>
      </div>
      <div className="max-h-[80vh] w-full overflow-x-auto">
        <table className="table table-xs table-zebra table-pin-rows table-pin-cols overflow-x-auto">
          <thead>
            <tr>
              <th></th>
              <td>Username</td>
              <td>Email</td>
              <td>Avatar</td>
              <td>Name</td>
              <td>Permissions</td>
              <td>Created At</td>
              <td>Last Updated At</td>
              <td>Deleted At</td>
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
                    className="w-6 h-6 rounded-full"
                    src={`https://www.gravatar.com/avatar/${md5(
                      user.email
                    )}?s=2048`}
                    alt=""
                  />
                </td>
                <td>{user.firstName + " " + user.lastName}</td>
                <td>{user.permissions.join(", ")}</td>
                <td>{new Date(user.createdAt * 1000).toLocaleString()}</td>
                <td>{new Date(user.updatedAt * 1000).toLocaleString()}</td>
                <td>
                  {user.deletedAt !== 0
                    ? new Date(user.deletedAt * 1000).toLocaleString()
                    : ""}
                </td>
                <td className="flex justify-center gap-1">
                  <button
                    onClick={() =>
                      (
                        document.getElementById(
                          `user_edit_modal_${user.id}`
                        ) as HTMLDialogElement
                      )?.showModal()
                    }
                    className="btn btn-xs btn-accent"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() =>
                      (
                        document.getElementById(
                          `user_delete_modal_${user.id}`
                        ) as HTMLDialogElement
                      )?.showModal()
                    }
                    className={clx(
                      "btn btn-xs",
                      user.deletedAt !== 0 ? "btn-error" : "btn-warning"
                    )}
                  >
                    Delete
                  </button>
                </td>
                <th>{i + 1}</th>
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
              <td>Created At</td>
              <td>Last Updated At</td>
              <td>Deleted At</td>
              <td>Actions</td>
              <th></th>
            </tr>
          </tfoot>
        </table>
      </div>
      {users?.map((user: any, i: number) => (
        <>
          <EditUserForm key={i} user={user} refetchUsers={refetchUsers} />
          <DeleteUserModal key={i} user={user} refetchUsers={refetchUsers} />
        </>
      ))}
    </div>
  );
};

export default ManageUsers;
