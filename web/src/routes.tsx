import { RouteObject, createBrowserRouter } from 'react-router-dom';
import { Root } from './routes/root';
import { Gallery } from './routes/gallery';

interface DOMRouterOpts {
  basename?: string;
  future?: Partial<Omit<RouterFutureConfig, "v7_prependBasename">>;
  hydrationData?: HydrationState;
  unstable_dataStrategy?: unstable_DataStrategyFunction;
  window?: Window;
}

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
  },
  {
    path: "/gallery",
    element: <Gallery />,
  },
] as RouteObject[],
{

} as DOMRouterOpts);