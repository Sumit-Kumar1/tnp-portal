import { AddEdit } from 'components/jobs/AddEdit';
import jobsService from 'services/jobs';

export default AddEdit;

export async function getServerSideProps({ params }) {
    const user = await jobsService.getById(params.id);
    return {
        props: { user }
    }
}