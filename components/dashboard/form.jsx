import { userService, alertService } from "services";
import { useState, useEffect } from "react";
import { useRouter } from "next/router";

export default FormData;

function FormData() {
  const [user, setUser] = useState(null);
  const router = useRouter();

  // fetch user and set default form values if in edit mode
  useEffect(() => {
    const subscription = userService.user.subscribe((x) => setUser(x));
    return () => subscription.unsubscribe();
  }, []);

  function onSubmit(e) {
    e.preventDefault();
    let accounts = [];
    for (const i = 0; i < 4; ++i) {
      accounts.push(e.target[i].value);
    }
    user.accounts = accounts;
    console.log(user);
    return userService
      .update(user.ErNo, user)
      .then(() => {
        alertService.success("User updated", { keepAfterRouteChange: true });
        router.push("/");
      })
      .catch(alertService.error);
  }

  return (
    <form
      onSubmit={onSubmit}
      className="flex flex-col border-2 rounded px-3 py-2 w-[1/2] md:w-[1/3]"
    >
      <label>Enter LeetCode userName</label>
      <input type="text" className="border-2 m-2 px-3"></input>
      <label>Enter interviewBit userName</label>
      <input type="text" className="border-2 m-2 px-3"></input>
      <label>Enter Codechef userName</label>
      <input type="text" className="border-2 m-2 px-3"></input>
      <label>Enter HackerRank userName</label>
      <input type="text" className="border-2 m-2 px-3"></input>
      <button type="submit" className="btn">
        submit
      </button>
    </form>
  );
}
