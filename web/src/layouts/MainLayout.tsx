import { Outlet } from "react-router-dom";
import { Theme } from "@radix-ui/themes";
import Nav from "@/components/ui/navbar";
// import Navbar from '../components/Navbar/Navbar';

const MainLayout = () => (
  <>
    {/* <Theme> */}
    <Nav>
      <Outlet />
    </Nav>
    {/* </Theme> */}
  </>
);

export default MainLayout;
