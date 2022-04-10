import { useRouter } from "next/router";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as Yup from "yup";

import { Link } from "components";
import { Layout } from "components/account";
import { userService, alertService } from "services";

export default Login;

function Login() {
  const router = useRouter();

  // form validation rules
  const validationSchema = Yup.object().shape({
    username: Yup.string().required("Username is required"),
    password: Yup.string().required("Password is required"),
  });
  const formOptions = { resolver: yupResolver(validationSchema) };

  // get functions to build form with useForm() hook
  const { register, handleSubmit, formState } = useForm(formOptions);
  const { errors } = formState;

  function onSubmit({ username, password }) {
    return userService
      .login(username, password)
      .then(() => {
        // get return url from query parameters or default to '/'
        const returnUrl = router.query.returnUrl || "/";
        router.push(returnUrl);
      })
      .catch(alertService.error);
  }

  return (
    <Layout>
      <div className="flex flex-col h-screen justify-center mx-auto gap-8">
        <h4 className="text-center text-2xl font-bold">Login</h4>
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="flex flex-col gap-1">
            <label className="font-bold">Username</label>
            <input
              name="username"
              type="text"
              {...register("username")}
              className={`border-2 rounded-md p-2
                 ${errors.username ? "is-invalid" : ""}`}
            />
            <div className="text-red-500">{errors.username?.message}</div>
            <label className="font-bold">Password</label>
            <input
              name="password"
              type="password"
              {...register("password")}
              className={`border-2 rounded-md p-2  ${
                errors.password ? "is-invalid" : ""
              }`}
            />
            <div className="text-red-500">{errors.password?.message}</div>
          </div>
          <div className="flex flex-col text-center">
            <button disabled={formState.isSubmitting} className="btn">
              {formState.isSubmitting && <span className=""></span>}
              Login
            </button>
            <Link href="/account/register" className="btn">
              Register
            </Link>
          </div>
        </form>
      </div>
    </Layout>
  );
}
