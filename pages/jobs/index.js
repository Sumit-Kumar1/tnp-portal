import { useState, useEffect } from "react";

import { Link } from "components";
import { jobService} from "services";

export default Index;

function Index() {
  const [jobs, setJobs] = useState(null);

  useEffect(() => {
    jobService.getAll().then((x) => setJobs(x));
  }, []);

  function deleteJob(id) {
    setJobs(
      jobs.map((x) => {
        if (x.id === id) {
          x.isDeleting = true;
        }
        return x;
      })
    );
    jobService.delete(id).then(() => {
      setJobs((jobs) => jobs.filter((x) => x.id !== id));
    });
  }

  return (
    <div>
      <h1 className="text-xl text-center font-bold mt-3">JOBS</h1>
      <Link href="/jobs/add" className="inline-block bg-transparent hover:bg-indigo-500 text-indigo-700 font-semibold hover:text-white py-2 px-4 border border-indigo-500 hover:border-transparent rounded mx-auto">
        Add Job
      </Link>
      <table className="table table-striped">
        <thead>
          <tr>
            <th style={{ width: "30%" }}>Name</th>
            <th style={{ width: "30%" }}>Email</th>
            <th style={{ width: "30%" }}>Role</th>
            <th style={{ width: "10%" }}>#</th>
          </tr>
        </thead>
        <tbody>
          {jobs &&
            jobs.map((job) => (
              <tr key={job.id}>
                <td>
                  {job.title} {job.description} {job.company}
                </td>
                <td>{job.email}</td>
                <td>{job.role}</td>
                <td style={{ whiteSpace: "nowrap" }}>
                  <Link
                    href={`/jobs/edit/${job.id}`}
                    className="mr-1 btn btn-sm btn-primary"
                  >
                    Edit
                  </Link>
                  <button
                    onClick={() => deleteJob(job.id)}
                    className="btn btn-sm btn-danger btn-delete-job"
                    disabled={job.isDeleting}
                  >
                    {job.isDeleting ? (
                      <span className="spinner-border spinner-border-sm"></span>
                    ) : (
                      <span>Delete</span>
                    )}
                  </button>
                </td>
              </tr>
            ))}
          {!jobs && (
            <tr>
              <td colSpan="4" className="text-center">
                <div className="spinner-border spinner-border-lg align-center"></div>
              </td>
            </tr>
          )}
          {jobs && !jobs.length && (
            <tr>
              <td colSpan="4" className="text-center">
                <div className="p-2">No jobs To Display</div>
              </td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
}
