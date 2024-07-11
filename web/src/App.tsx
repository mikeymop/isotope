import { RouterProvider } from "react-router-dom";
import { Theme } from '@radix-ui/themes';

import { router } from './routes'

const App = () => (
  <RouterProvider router={router} />
);

export default App
