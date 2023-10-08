import { zodResolver } from "@hookform/resolvers/zod";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { useForm } from "react-hook-form";
import z from "zod";
import CloseIconSvg from "../../../SVGs/CloseIconSvg";
import { handleError } from "../../../utils/error.utils";
import { promiseToast } from "../../../utils/toast.utils";

const EditUserForm = ({
  user,
  refetchUsers,
}: {
  user: any;
  refetchUsers: () => void;
}) => {
  const { data: allPermissions } = useQuery({
    queryKey: ["all-permissions"],
    queryFn: () =>
      axios.get("/api/auth/permissions").then((res) => res.data as string[]),
  });

  const { register, watch, setValue, handleSubmit } = useForm<IFormData>({
    resolver: zodResolver(editUserFormSchema),
    values: user,
  });
  const selectedPermissions = watch("permissions");

  const handleUpdateUser = (data: IFormData) =>
    promiseToast(axios.patch(`/api/users/${user.id}`, data), {
      loading: "Updating...",
      success: "Successfully updated user!",
    })
      .then(() =>
        (
          document.getElementById(
            `user_edit_modal_${user.id}`
          ) as HTMLDialogElement
        ).close()
      )
      .then(refetchUsers)
      .catch(handleError);

  return (
    <dialog id={`user_edit_modal_${user.id}`} className="modal relative">
      <div className="modal-box">
        <h3 className="font-bold text-lg">Edit User</h3>
        <form onSubmit={handleSubmit(handleUpdateUser)}>
          <div>
            <label className="label">
              <span className="label-text">Username</span>
            </label>
            <input
              className="input input-bordered w-full"
              type="text"
              {...register("username")}
            />
          </div>
          <div>
            <label className="label">
              <span className="label-text">Email</span>
            </label>
            <input
              className="input input-bordered w-full"
              type="text"
              {...register("email")}
            />
          </div>
          <div>
            <label className="label">
              <span className="label-text">Password</span>
            </label>
            <input
              className="input input-bordered w-full"
              type="password"
              {...register("password")}
            />
          </div>
          <div>
            <label className="label">
              <span className="label-text">First Name</span>
            </label>
            <input
              className="input input-bordered w-full"
              type="text"
              {...register("firstName")}
            />
          </div>
          <div>
            <label className="label">
              <span className="label-text">Last Name</span>
            </label>
            <input
              className="input input-bordered w-full"
              type="text"
              {...register("lastName")}
            />
          </div>
          <div>
            <label className="label">
              <span className="label-text">Permissions</span>
            </label>
            <div className="flex flex-wrap gap-3 items-center">
              {allPermissions?.map((permission) => (
                <div
                  key={permission}
                  className="form-control w-48 bg-base-300 rounded p-1 text-sm"
                >
                  <label className="label cursor-pointer">
                    <span className="label-text">{permission}</span>
                    <input
                      type="checkbox"
                      className="checkbox checkbox-xs"
                      checked={
                        selectedPermissions?.includes(permission) || false
                      }
                      onChange={(e) =>
                        setValue(
                          "permissions",
                          Array.from(
                            new Set(
                              e.target.checked
                                ? [...selectedPermissions, permission]
                                : selectedPermissions?.filter(
                                    (p) => p !== permission
                                  )
                            )
                          )
                        )
                      }
                    />
                  </label>
                </div>
              ))}
            </div>
          </div>
          <button type="submit" className="btn btn-primary my-8">
            Update
          </button>
        </form>
        <div className="modal-action absolute top-0 right-3">
          <form method="dialog">
            <button className="btn btn-xs">
              <CloseIconSvg />
            </button>
          </form>
        </div>
      </div>
    </dialog>
  );
};

export default EditUserForm;

const editUserFormSchema = z.object({
  username: z.string(),
  password: z.string(),
  email: z.string().email(),
  firstName: z.string(),
  lastName: z.string(),
  permissions: z.string().array(),
});

type IFormData = z.infer<typeof editUserFormSchema>;
