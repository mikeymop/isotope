import { Aperture } from "lucide-react";
import { Link } from "react-router-dom";

/** Isotope Navbar
 * @see https://lucide.dev/icons/
 */
const Nav = () => {
  return (
    <header className="sticky top-0 flex h-12 items-center gap-4 border-b bg-background px-4 md:px-6">
      <nav className="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
        <Link
          to="/"
          className="flex items-center gap-2 text-lg font-semibold md:text-base"
        >
          <Aperture className="h-6 w-6" />
          <span className="sr-only">Acme Inc</span>
        </Link>
        <Link
          to="/"
          className="text-foreground transition-colors hover:text-foreground"
        >
          Dashboard
        </Link>
        <Link
          to="/gallery"
          className="text-muted-foreground transition-colors hover:text-foreground"
        >
          Galleries
        </Link>
        <Link
          to="#"
          className="text-muted-foreground transition-colors hover:text-foreground"
        >
          Photos
        </Link>
        <Link
          to="#"
          className="text-muted-foreground transition-colors hover:text-foreground"
        >
          Customers
        </Link>
      </nav>
    </header>
  );
};

export default Nav;
