import React, { useRef, useState } from 'react';

const Sidebar = () => {
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
                <App_template app={app} />
            ))}
        </ul>
        )}


      </div>
    </div>
  );
};

export default Sidebar;


type Application = {
  id: number;
  name: string;
  url: string;
};



const App_template = ({ app }: { app: Application }) => {
  return (
    <li key={app.id}>
      <p>{app.name}</p>
      <p>{app.url}</p>
    </li>
  );
};