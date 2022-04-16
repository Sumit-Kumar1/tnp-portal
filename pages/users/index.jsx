import { useState, useEffect } from "react";

import { Link, Spinner } from "components";
import { Layout } from "components/users";
import { userService } from "services";

export default Index;

function Index() {
  const [users, setUsers] = useState(null);

  useEffect(() => {
    userService.getAll().then((x) => setUsers(x));
  }, []);

  function deleteUser(ErNo) {
    setUsers(
      users.map((x) => {
        if (x.ErNo === ErNo) {
          x.isDeleting = true;
        }
        return x;
      })
    );
    userService.delete(ErNo).then(() => {
      setUsers((users) => users.filter((x) => x.ErNo !== ErNo));
    });
  }

  return (
    <Layout>
      <Link
        href="/users/add"
        className="border-2 rounded-md px-2 py-2 block my-2 w-1/4 mx-auto text-center text-white bg-indigo-700 font-bold hover:bg-indigo-900"
      >
        Add User
      </Link>
      <div className="p-5">
        <table className="min-w-full text-center">
          <thead className="border-b bg-gray-800 text-white">
            <tr>
              <th className="text-sm font-medium text-white px-6 py-4">
                Enrollment Number
              </th>
              <th className="text-sm font-medium text-white px-6 py-4">
                First Name
              </th>
              <th className="text-sm font-medium text-white px-6 py-4">
                Last Name
              </th>
              <th className="text-sm font-medium text-white px-6 py-4">#</th>
            </tr>
          </thead>
          <tbody className=" ">
            {users &&
              users.map((user) => (
                <tr key={user.ErNo} className="border-b hover:bg-gray-200">
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {user.ErNo}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {user.firstName}
                  </td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {user.lastName}
                  </td>
                  <td>
                    <div className="flex flex-row gap-2">
                      <Link
                        href={`/users/edit/${user.ErNo}`}
                        className="btn mx-auto my-1"
                      >
                        Edit
                      </Link>
                      <button
                        onClick={() => deleteUser(user.ErNo)}
                        className="btn-del my-1 mx-auto"
                        disabled={user.isDeleting}
                      >
                        {user.isDeleting ? (
                          <span className="spinner-border spinner-border-sm"></span>
                        ) : (
                          <span>Delete</span>
                        )}
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
            {!users && (
              <tr>
                <td colSpan="4">
                  <Spinner />
                </td>
              </tr>
            )}
            {users && !users.length && (
              <tr>
                <td colSpan="4" className="text-center">
                  <div className="p-2">No Users To Display</div>
                </td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </Layout>
  );
}
