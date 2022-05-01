import NextLink from "next/link";

export { JobCard };
function JobCard({ id, ...details }) {
  return (
    <div className="max-w-sm p-2 overflow-hidden rounded shadow-lg hover:shadow-xl hover:border-4 hover:border-indigo-400 hover:p-4">
      <div className="px-6 py-4">
        <h2 className="mb-2 text-xl font-bold">{details.title}</h2>
        <p className="text-base text-gray-700">
          Description: {details.description}
        </p>
      </div>
      <div>
        <span className="inline-block px-3 py-1 mb-2 mr-2 text-sm font-semibold text-gray-700 bg-gray-200 rounded-full">
          {details.role}
        </span>
        <br />
        <span className="inline-block px-3 py-1 mb-2 mr-2 text-sm font-semibold text-gray-700 bg-gray-200 rounded-full">
          Start: {details.startDate}
        </span>{" "}
        <span className="inline-block px-3 py-1 mb-2 mr-2 text-sm font-semibold text-gray-700 bg-gray-200 rounded-full">
          End : {details.endDate}
        </span>
      </div>

      <NextLink href={`/jobs/[id]`} as={`/jobs/${id}`} passHref>
        <a className="btn">More Info</a>
      </NextLink>
    </div>
  );
}
