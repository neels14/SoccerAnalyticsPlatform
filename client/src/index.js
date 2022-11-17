import React from 'react';
import ReactDOM from 'react-dom/client';
import {
  createBrowserRouter,
  RouterProvider
} from "react-router-dom";
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import './index.css'
import Root from './routes/Root';
import Home from './routes/Home'
import WorldCups from './routes/WorldCups';
import Countries from './routes/Countries';
import Players from './routes/Players';
import Matches from './routes/Matches';
import Feature1 from './routes/Feature1';
import Feature2 from './routes/Feature2';
import Feature3 from './routes/Feature3';
import Feature4 from './routes/Feature4';
import Feature5 from './routes/Feature5';
import Feature6 from './routes/Feature6';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root/>,
    children: [
      {
        path: "",
        element: <Home/>,
      },
      {
        path: "worldcups",
        element: <WorldCups/>,
      },
      {
        path: "countries",
        element: <Countries/>,
      },
      {
        path: "players",
        element: <Players/>,
      },
      {
        path: "matches",
        element: <Matches/>,
      },
      {
        path: "features/1",
        element: <Feature1/>,
      },
      {
        path: "features/2",
        element: <Feature2/>,
      },
      {
        path: "features/3",
        element: <Feature3/>,
      },
      {
        path: "features/4",
        element: <Feature4/>,
      },
      {
        path: "features/5",
        element: <Feature5/>,
      },
      {
        path: "features/6",
        element: <Feature6/>,
      },
    ]
  }
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
