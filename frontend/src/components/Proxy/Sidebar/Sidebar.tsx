import React, { useRef, useState } from 'react';
import type { Application } from '../Proxy';
import type { SidebarClasses } from './Sidebar_classes';


type SidebarProps = {
  set_current_app: React.Dispatch<React.SetStateAction<Application | null>>;
  classes: SidebarClasses;
};

const Sidebar = ({ set_current_app, classes }: SidebarProps) => {
  let appNameRef = useRef<HTMLInputElement>(null);
  let appUrlRef = useRef<HTMLInputElement>(null);
  let [isDisabled, setIsDisabled] = useState<boolean>(false);

  let HandleAdd = () => {
    let name = appNameRef.current?.value;
    let url = appUrlRef.current?.value;
    if (!name || !url) return;

    setIsDisabled(true);

    // Add your logic here

    setIsDisabled(false);
   };


   let [applications,setapp]=useState<Application[]>([])

  return (
    <div className="sidebar_container">
      <div className="sidebar_background">
        <div>
            <input ref={appNameRef} type="text" placeholder="Enter App Name" />
            <input ref={appUrlRef} type="text" placeholder="Enter App Url" />
            <button onClick={HandleAdd} disabled={isDisabled}>Add</button>
        </div>


        {applications.length > 0 && (
        <ul>
            {applications.map((app) => (
                <App_template app={app} set_current_app={set_current_app}/>
            ))}
        </ul>
        )}


      </div>
    </div>
  );
};

export default Sidebar;



const App_template = ({app,set_current_app,}: {
  app: Application;
  set_current_app: React.Dispatch<React.SetStateAction<Application | null>>;
}) => {
  return (
    <li key={app.id} onClick={() => set_current_app(app)}>
      <p>{app.name}</p>
      <p>{app.url}</p>
    </li>
  );
};