import React, { useEffect, useRef, useState } from 'react';
import type { Application } from '../Proxy';
import type { SidebarClasses } from './Sidebar_classes';
import { api } from '../../../Api/api';


type SidebarProps = {
  set_current_app: React.Dispatch<React.SetStateAction<Application | null>>;
  classes: SidebarClasses;
};

const Sidebar = ({ set_current_app, classes }: SidebarProps) => {
  let appNameRef = useRef<HTMLInputElement>(null);
  let appUrlRef = useRef<HTMLInputElement>(null);
  let [isDisabled, setIsDisabled] = useState<boolean>(false);

  let HandleAdd = async () => {
    let name = appNameRef.current?.value;
    let url = appUrlRef.current?.value;
    if (!name || !url) return;

    setIsDisabled(true);

    // Add your logic here
    let res=await api.proxy.CreateApp(name,url)

    if (!res.success) {
      setIsDisabled(false);
      return
    }

    let app: Application = {
        id: res.response.data.id,
        name: res.response.data.name,
        url: url
    };

    setapp(prev => [...prev, app]);

    setIsDisabled(false);
  };


  let [applications,setapp]=useState<Application[]>([])

  useEffect(() => {

    let a=async()=>{
      let res=await api.proxy.GetApps()

      if (!res.success){
        return
      }

      let apps = res.response.data as Application[];
      
      setapp(apps)
    }
    
    a()
    
  }, [])
  


  return (
    <div className="sidebar_container">
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



        <p>Active LSiteners</p>
    </div>
  );
};

export default Sidebar;



const App_template = ({app,set_current_app,}: {
  app: Application;
  set_current_app: React.Dispatch<React.SetStateAction<Application | null>>;
}) => {

  let nameRef = useRef<HTMLInputElement>(null);
  let urlRef = useRef<HTMLInputElement>(null);

  let handle_update=async()=>{
    let newName = nameRef.current?.value ?? "";
    let newUrl = urlRef.current?.value || app.url;

    let res = await api.proxy.UpdateApp(app.name, newName, newUrl);

    if (!res.success) {
        return;
    }
  }

  let handle_delete = async () => {
      const res = await api.proxy.DeleteApp(app.id);

      if (!res.success) {
          return;
      }
  };

  return (
    <li key={app.id} onClick={() => set_current_app(app)}>
      <p>Name {app.name}</p>
      <p>url {app.url}</p>

      <input ref={nameRef} type='text' placeholder='New name'/>
      <input ref={urlRef} type ='text' placeholder="New url"/>
      <button onClick={handle_update}>update</button>
      <button onClick={handle_delete}>delete</button>
    </li>
  );
};


