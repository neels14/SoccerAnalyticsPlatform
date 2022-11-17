import { Outlet } from "react-router-dom";
import Navbar from './Navbar';

function Root() {
    return (
        <>
            <Navbar/>
            <div style={{margin: 8}}>
                <Outlet/>
            </div>
        </>
    );
  }
  
  export default Root;
  