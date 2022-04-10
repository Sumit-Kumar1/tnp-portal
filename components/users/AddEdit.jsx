import { useRouter } from "next/router";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as Yup from "yup";

import { Link } from "components";
import { userService, alertService } from "services";

export { AddEdit };

function AddEdit(props) {
  const user = props?.user;
  const isAddMode = !user;
  const router = useRouter();

  // form validation rules
  const validationSchema = Yup.object().shape({
    firstName: Yup.string().required("*Required"),
    lastName: Yup.string().required("*Required"),
    ErNo: Yup.number()
      .test("len", "Must be exactly 6 characters", (val) => {
        if (val) return val.toString().length === 6;
      })
      .required("Enrollment Number is required")
      .typeError("you must specify a number")
      .min(181000, "Enter valid Enrollment Number")
      .max(219999, "enter valid Enrollment number"),
    password: Yup.string()
      .transform((x) => (x === "" ? undefined : x))
      .concat(isAddMode ? Yup.string().required("Password is required") : null)
      .min(6, "Password must be at least 6 characters"),
  });
  const formOptions = { resolver: yupResolver(validationSchema) };

  // set default form values if in edit mode
  if (!isAddMode) {
    formOptions.defaultValues = props.user;
  }

  // get functions to build form with useForm() hook
  const { register, handleSubmit, reset, formState } = useForm(formOptions);
  const { errors } = formState;

  function onSubmit(data) {
    return isAddMode ? createUser(data) : updateUser(user.ErNo, data);
  }

  function createUser(data) {
    return userService
      .register(data)
      .then(() => {
        alertService.success("User added", { keepAfterRouteChange: true });
        router.push(".");
      })
      .catch(alertService.error);
  }

  function updateUser(ErNo, data) {
    return userService
      .update(ErNo, data)
      .then(() => {
        alertService.success("User updated", { keepAfterRouteChange: true });
        router.push("..");
      })
      .catch(alertService.error);
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="">
        <div className="form-group col">
          <label>First Name</label>
          <input
            name="firstName"
            type="text"
            {...register("firstName")}
            className={` ${errors.firstName ? "is-invalid" : ""}`}
          />
          <div className="">{errors.firstName?.message}</div>
        </div>
        <div className="form-group col">
          <label>Last Name</label>
          <input
            name="lastName"
            type="text"
            {...register("lastName")}
            className={` ${errors.lastName ? "is-invalid" : ""}`}
          />
          <div className="">{errors.lastName?.message}</div>
        </div>
      </div>
      <div className="">
        <div className="form-group col">
          <label>Enrollment Number</label>
          <input
            name="ErNo"
            type="text"
            {...register("ErNo")}
            className={` ${errors.ErNo ? "is-invalid" : ""}`}
          />
          <div className="">{errors.email?.message}</div>
        </div>
        <div className="form-group col">
          <label>
            Password
            {!isAddMode && (
              <em className="ml-1">(Leave blank to keep the same password)</em>
            )}
          </label>
          <input
            name="password"
            type="password"
            {...register("password")}
            className={` ${errors.password ? "is-invalid" : ""}`}
          />
          <div className="">{errors.password?.message}</div>
        </div>
      </div>
      <div className="form-group">
        <button
          type="submit"
          disabled={formState.isSubmitting}
          className="btn btn-primary mr-2"
        >
          {formState.isSubmitting && (
            <span className="spinner-border spinner-border-sm mr-1"></span>
          )}
          Save
        </button>
        <button
          onClick={() => reset(formOptions.defaultValues)}
          type="button"
          disabled={formState.isSubmitting}
          className="btn btn-secondary"
        >
          Reset
        </button>
        <Link href="/users" className="btn btn-link">
          Cancel
        </Link>
      </div>
    </form>
  );
}
