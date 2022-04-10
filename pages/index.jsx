import { userService } from 'services';
import { Link } from 'components';

export default Home;

function Home() {
    return (
        <div className="p-4">
            <div className="container">
                <h1 className ="text-xl font-bold">Hi {userService.userValue?.firstName}!</h1>
                <p>You&apos;re logged in with Next.js & JWT!!</p>
                <p><Link href="/users" className="border-2 border-blue-600 bg-indigo-500 rounded px-3 py-2 inline-block text-white hover:bg-indigo-700">Manage Users</Link></p>
            </div>
        </div>
    );
}
