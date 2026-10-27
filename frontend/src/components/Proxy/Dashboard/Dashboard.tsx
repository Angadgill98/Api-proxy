import type { Application } from '../Proxy';
import type { DashboardClasses } from './Dashboard_classes';

interface DashboardProps {
  app: Application | null;
  classes: DashboardClasses;
}

const Dashboard = ({ app, classes }: DashboardProps) => {
  return (
    <div>
      {app ? app.name : "No application selected"}
    </div>
  );
};

export default Dashboard;



const Options=()=>{

}