import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React from "react";
import ReactDOM from "react-dom/client";
import { Toaster } from "react-hot-toast";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import "./index.css";
import DashboardLayout from "./layouts/dashboard-layout";
import RootLayout from "./layouts/root";
import SettingsLayout from "./layouts/settings-layout";
import LoginPage from "./pages/login";
import ManageUsers from "./pages/settings/manage-users";

const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      {
        path: "login",
        element: <LoginPage />,
      },
      {
        path: "dashboard",
        element: <DashboardLayout />,
        children: [],
      },
      {
        path: "settings",
        element: <SettingsLayout />,
        children: [
          {
            path: "manage-users",
            element: <ManageUsers />,
          },
        ],
      },
    ],
  },
]);

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
      <Toaster position="top-right" />
    </QueryClientProvider>
  </React.StrictMode>
);
