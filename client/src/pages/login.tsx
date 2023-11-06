import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import z from "zod";
import { handleError } from "../utils/error.utils";

const formSchema = z.object({
  usernameOrEmail: z.string().min(1).max(30),
  password: z.string().min(1).max(80),
});

type IFormData = z.infer<typeof formSchema>;

const LoginPage = () => {
  const { register, handleSubmit } = useForm<IFormData>({
    resolver: zodResolver(formSchema),
  });
  const navigate = useNavigate();
  const handleLogin = (data: IFormData) =>
    axios
      .post("/api/auth/login", data)
      .then(() => navigate("/dashboard"))
      .catch(handleError);
  return (
    <div>
      <form onSubmit={handleSubmit(handleLogin)} className="max-w-md p-8">
        <div>
          <label className="label">
            <span className="label-text">Username or Email</span>
          </label>
          <input
            className="input input-bordered w-full"
            type="text"
            {...register("usernameOrEmail")}
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
        <button className="btn btn-sm my-4" type="submit">
          Login
        </button>
      </form>
    </div>
  );
};

export default LoginPage;
