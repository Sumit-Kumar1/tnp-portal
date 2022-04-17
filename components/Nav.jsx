import { useState, useEffect } from "react";

import { NavLink } from ".";
import { userService } from "services";

export { Nav };

function Nav() {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const subscription = userService.user.subscribe((x) => setUser(x));
    return () => subscription.unsubscribe();
  }, []);

  function logout() {
    userService.logout();
  }

  // only show nav when logged in
  if (!user) return null;

  return (
    <nav>
      <div className="flex flex-row gap-5 text-xl bg-black text-white p-5">
        {user.role === "admin" ? (
          <NavLink href="/admin" exact className="">
            Home
          </NavLink>
        ) : (
          <NavLink href="/student" exact className="">
            Home
          </NavLink>
        )}

        {userService.userValue.role === "admin" && (
          <NavLink href="/users" className="">
            Users
          </NavLink>
        )}
        <a onClick={logout} className="">
          Logout
        </a>
      </div>
    </nav>
  );
}
