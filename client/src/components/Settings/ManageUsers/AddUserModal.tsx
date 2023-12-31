import { zodResolver } from "@hookform/resolvers/zod";
import { useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useForm } from "react-hook-form";
import z from "zod";
import CloseIconSvg from "../../../SVGs/CloseIconSvg";
import { handleError } from "../../../utils/error.utils";
import { promiseToast } from "../../../utils/toast.utils";

const formSchema = z.object({
  username: z.string().min(1, "Required"),
  password: z.string().min(1, "Required"),
  email: z.string().email().min(1, "Required"),
  firstName: z.string().min(1, "Required"),
  lastName: z.string().min(1, "Required"),
});

type IFormData = z.infer<typeof formSchema>;

const AddUserModal = () => {
  const queryClient = useQueryClient();
  const { register, handleSubmit } = useForm<IFormData>({
    resolver: zodResolver(formSchema),
  });
  const handleAddUser = (data: IFormData) =>
    promiseToast(
      axios
        .post("/api/users", data)
        .then(() =>
          (
            document.getElementById(`user_add_modal`) as HTMLDialogElement
          )?.close()
        )
        .then(() => queryClient.invalidateQueries(["all-users"]))
        .catch(handleError)
    );
  return (
    <dialog id={`user_add_modal`} className="modal">
      <div className="modal-box">
        <h3 className="font-bold text-lg mt-6">Add New User</h3>
        <form onSubmit={handleSubmit(handleAddUser)}>
          <div>
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
                type="text"
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
          </div>
          <button type="submit" className="btn btn-sm btn-primary mt-3">
            Add User
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

export default AddUserModal;
