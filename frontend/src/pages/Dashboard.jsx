import { useEffect, useState } from "react";
import StatsCard from "../components/StatsCard";

function Dashboard() {
  const [stats, setStats] = useState({});
  const [users, setUsers] = useState([]);

  useEffect(() => {
    // Fetch dashboard stats
    fetch("/admin/stats")
      .then((res) => res.json())
      .then((data) => setStats(data));

    // Fetch user list
    fetch("/admin/users")
      .then((res) => res.json())
      .then((data) => setUsers(data));
  }, []);

  return (
    <div className="p-6 font-sans">
      <h1 className="text-3xl font-bold mb-6">Admin Dashboard</h1>

      <div className="grid grid-cols-2 gap-4 mb-8">
        <StatsCard title="Total Users" value={stats.totalUsers} />
        <StatsCard title="Active Users" value={stats.activeUsers} />
      </div>

      <h2 className="text-xl font-semibold">User List</h2>
      <ul className="list-disc pl-5 mt-2">
        {users.map((user) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
}

export default Dashboard;
