import { Aperture } from "lucide-react";
import { Link, useLocation } from "react-router-dom";
import { cn } from "@/lib/utils";

/** Renders a Navbar button. */
const NavLink = ({ to, title }: { to: string; title: string }) => {
  const { pathname } = useLocation();
  return (
    <Link
      to={to}
      className={cn(
        "transition-colors hover:text-foreground",
        pathname === to ? "font-semibold" : "text-foreground",
      )}
    >
      {title}
    </Link>
  );
};

/** Isotope Navbar
 * @see https://lucide.dev/icons/
 */
const Nav = (): JSX.Element => {
  return (
    <header className="sticky top-0 flex h-12 items-center gap-4 border-b bg-background px-4 md:px-6">
      <nav className="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
        <Link to="/" className="flex items-center gap-2 text-lg md:text-base">
          <Aperture className="h-6 w-6" />
          <span className="sr-only">Isotope</span>
        </Link>
        <NavLink to="/" title="Dashboard" />
        <NavLink to="/gallery" title="Galleries" />
        <NavLink to="/photos" title="Photos" />
        <NavLink to="/customers" title="Customers" />
      </nav>
    </header>
  );
};

export default Nav;
