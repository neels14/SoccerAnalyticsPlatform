import { Outlet, useMatch } from "react-router-dom";
import Navbar from './Navbar';
import { useState } from 'react';
import UserContext from "./shared/user-context";
import Login from "./Login";

function Root() {
    const [user, setUser] = useState("")
    const isSignUp = useMatch({ path: `/signup`, end: true})

    return (
        <UserContext.Provider value={{ user, setUser }}>
            <Navbar/>
            { user === "" && !isSignUp?
                <Login/> :
                <div style={{margin: 8}}>
                    <Outlet/>
                </div>
            }
        </UserContext.Provider>
    );
  }
  
  export default Root;
  