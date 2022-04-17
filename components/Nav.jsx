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
        <NavLink href="/" exact className="">
          Home
        </NavLink>

        {userService.userValue.role == "admin" && (
          <NavLink href="/users" className="">
            Users
          </NavLink>
        )}
        <NavLink href="/notification" className="">
          notifications
        </NavLink>
        <a onClick={logout} className="">
          Logout
        </a>
        <h1 className="text-xl font-bold text-green-600">
          Hi {userService.userValue?.firstName}! ({userService.userValue?.ErNo})
        </h1>
      </div>
    </nav>
  );
}
