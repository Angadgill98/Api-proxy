import { api } from '../../../Api/api';
import type { Application, Listener } from '../Proxy';
import type { DashboardClasses } from './Dashboard_classes';

interface DashboardProps {
  app: Application | null;
  classes: DashboardClasses;
  listeners: Listener[];
  setlisteners: React.Dispatch<React.SetStateAction<Listener[]>>;
}

const Dashboard = ({ app, classes ,listeners, setlisteners}: DashboardProps) => {



  let handle_listners_add=async ()=>{
    if (!app ) return;

    const res = await api.proxy.AddListener(app.name);

    if (!res.success) return;

    const exists = listeners.some(listener => listener.app.id === app.id);

    if (!exists) {
        setlisteners(prev => [...prev, { app }]);
    }

  }


  let handle_remove_listners=async ()=>{
    if (!app ) return;
  
    const res = await api.proxy.RemoveListener(app.name);

    if (!res.success) return;

    setlisteners(prev => prev.filter(listener => listener.app.id !== app.id));
  }

  return (
    <div>
      <button onClick={()=>handle_listners_add()}>Add to listeners</button>
      <button onClick={()=>handle_remove_listners()}>Remove</button>
      {app?.name}
    </div>
  );
};

export default Dashboard;



const Options=()=>{

}