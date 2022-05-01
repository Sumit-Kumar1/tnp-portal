import { Layout } from "components/users";
import { AddEdit } from "components/jobs/AddEditJobs";

export default Add;

function Add() {
  return (
    <Layout>
      <h1 className="block my-5 text-xl font-bold text-center">Add Job</h1>{" "}
      <AddEdit />
    </Layout>
  );
}
