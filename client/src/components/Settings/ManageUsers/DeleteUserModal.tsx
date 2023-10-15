import axios from "axios";
import CloseIconSvg from "../../../SVGs/CloseIconSvg";
import { handleError } from "../../../utils/error.utils";
import { promiseToast } from "../../../utils/toast.utils";

const DeleteUserModal = ({
  user,
  refetchUsers,
}: {
  user: any;
  refetchUsers: () => void;
}) => {
  const handleDelete = () =>
    promiseToast(axios.delete(`/api/users/${user.id}`), {
      loading: "Deleting...",
      success: "Successfully deleted user!",
    })
      .then(() =>
        (
          document.getElementById(
            `user_delete_modal_${user.id}`
          ) as HTMLDialogElement
        ).close()
      )
      .then(refetchUsers)
      .catch(handleError);
  return (
    <dialog id={`user_delete_modal_${user.id}`} className="modal">
      <div className="modal-box">
        <h3 className="font-bold text-lg mt-6">
          Are you sure you want to delete the user
          {user.deletedAt !== 0 ? " permanently" : ""}?
        </h3>
        <div className="flex justify-end">
          <button onClick={handleDelete} className="btn btn-error btn-sm my-10">
            Delete
          </button>
        </div>
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

export default DeleteUserModal;
