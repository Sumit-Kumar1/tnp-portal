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
      <h1>Users</h1>
      <Link href="/users/add" className="btn btn-sm btn-success mb-2">
        Add User
      </Link>
      <table className="table-auto">
        <thead>
          <tr>
            <th>Enrollment Number</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th style={{ width: "10%" }}></th>
          </tr>
        </thead>
        <tbody>
          {users &&
            users.map((user) => (
              <tr key={user.ErNo}>
                <td>{user.ErNo}</td>
                <td>{user.firstName}</td>
                <td>{user.lastName}</td>
                <td style={{ whiteSpace: "nowrap" }}>
                  <Link href={`/users/edit/${user.ErNo}`} className="btn ">
                    Edit
                  </Link>
                  <button
                    onClick={() => deleteUser(user.ErNo)}
                    className="btn-del"
                    disabled={user.isDeleting}
                  >
                    {user.isDeleting ? (
                      <span className="spinner-border spinner-border-sm"></span>
                    ) : (
                      <span>Delete</span>
                    )}
                  </button>
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
    </Layout>
  );
}
