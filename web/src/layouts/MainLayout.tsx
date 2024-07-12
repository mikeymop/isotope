import Nav from "@/components/ui/navbar";

const MainLayout = ({ children }: { children: React.ReactElement }) => (
  <>
    <Nav />
    {children}
  </>
);

export default MainLayout;
