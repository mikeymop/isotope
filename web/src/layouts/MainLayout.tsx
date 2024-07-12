// import { Outlet } from "react-router-dom";
import Nav from "@/components/ui/navbar";

const MainLayout = ({ children }: { children: React.ReactElement }) => (
  <>
    {/* <Theme> */}
    <Nav />
    {children}
    {/* </Theme> */}
  </>
);

export default MainLayout;
