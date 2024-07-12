import { useState } from "react";
import reactLogo from "../../assets/react.svg";
import viteLogo from "/vite.svg";
import "./Root.css";
// import MainLayout from "../../layouts/MainLayout";

const Root = () => {
  const [count, setCount] = useState(0);

  return (
    <>
      <div className="flex flex-col m-auto content-center justify-center">
        <div className="flex flex-row content-center justify-center">
          <a href="https://vitejs.dev" target="_blank">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <div className="flex justify-center">
          <h1>Welcome to Isotope</h1>
        </div>
      </div>
    </>
  );
};

export default Root;
