import { AddEditJob } from 'components/jobs/AddEditJobs';
import jobsService from 'services/jobs';

export default AddEditJob;

export async function getServerSideProps({ params }) {
    const job = await jobsService.getById(params.id);
    return {
        props: { job }
    }
}