import type { Application } from '../Proxy';

const Dashboard = ({ app }: { app: Application | null }) => {
  return (
    <div>
      {app ? app.name : "No application selected"}
    </div>
  );
};

export default Dashboard;