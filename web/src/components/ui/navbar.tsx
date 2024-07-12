import { Aperture } from "lucide-react";
import { Link } from "react-router-dom";
import { cn } from "@/lib/utils";

/** Return tw bold font if path matches current route. */
const isActive = (path: string) =>
  path === window.location.pathname ? "font-semibold" : "text-foreground";

/** Isotope Navbar
 * @see https://lucide.dev/icons/
 */
const Nav = (): JSX.Element => {
  console.log("Active: ", window.location.pathname, isActive("/gallery"));
  return (
    <header className="sticky top-0 flex h-12 items-center gap-4 border-b bg-background px-4 md:px-6">
      <nav className="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
        <Link to="/" className="flex items-center gap-2 text-lg md:text-base">
          <Aperture className="h-6 w-6" />
          <span className="sr-only">Acme Inc</span>
        </Link>
        <Link
          to="/"
          className={cn(
            "transition-colors hover:text-foreground",
            isActive("/"),
          )}
        >
          Dashboard
        </Link>
        <Link
          to="/gallery"
          className={cn(
            "transition-colors hover:text-foreground",
            isActive("/gallery"),
          )}
        >
          Galleries
        </Link>
        <Link
          to="#"
          className={cn(
            "transition-colors hover:text-foreground",
            isActive("/photos"),
          )}
        >
          Photos
        </Link>
        <Link
          to="#"
          className={cn(
            "transition-colors hover:text-foreground",
            isActive("/customers"),
          )}
        >
          Customers
        </Link>
      </nav>
    </header>
  );
};

export default Nav;
