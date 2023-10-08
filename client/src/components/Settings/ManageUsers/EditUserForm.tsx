/* eslint-disable @typescript-eslint/no-explicit-any */
const EditUserForm = ({ user }: { user: any }) => {
  return (
    <dialog id={`user_edit_modal_${user.id}`} className="modal">
      <div className="modal-box">
        <h3 className="font-bold text-lg">Hello {user.id}!</h3>
        <p className="py-4">Press ESC key or click the button below to close</p>
        <div className="modal-action">
          <form method="dialog">
            <button className="btn">Close</button>
          </form>
        </div>
      </div>
    </dialog>
  );
};

export default EditUserForm;