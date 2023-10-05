import { useEffect, useState } from "react";
import { BiCog } from "react-icons/bi";
import { Outlet } from "react-router-dom";
import Dropdown from "../components/Dropdown";
import { themes } from "../constants/themes";

const RootLayout = () => {
  const [selectedTheme, setSelectedTheme] = useState<string | undefined>(
    undefined
  );
  useEffect(() => {
    if (selectedTheme)
      document
        .getElementsByTagName("html")[0]
        .setAttribute("data-theme", selectedTheme);
  }, [selectedTheme]);
  useEffect(() => {
    const theme = localStorage.getItem("theme");
    if (theme) setSelectedTheme(theme);
  }, []);
  const handleSelectTheme = (theme: string) => {
    setSelectedTheme(theme);
    localStorage.setItem("theme", theme);
  };
  return (
    <main>
      <div className="fixed top-0 right-0">
        <Dropdown
          items={themes.map((theme) => () => (
            <a key={theme} onClick={() => handleSelectTheme(theme)}>
              {theme.toUpperCase()}
            </a>
          ))}
          header="Themes"
        >
          <BiCog />
        </Dropdown>
      </div>
      <Outlet />
    </main>
  );
};

export default RootLayout;
