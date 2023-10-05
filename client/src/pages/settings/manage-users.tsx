import { useQuery } from "@tanstack/react-query";
import axios from "axios";

const ManageUsers = () => {
  const { data: users } = useQuery({
    queryKey: ["users"],
    queryFn: () => axios.get("/api/users").then((res) => res.data),
  });

  return (
    <div>
      <h1>Manage Users</h1>
      <div>
        <pre>{JSON.stringify(users, null, 2)}</pre>
      </div>
    </div>
  );
};

export default ManageUsers;
