import React from "react";

const Dropdown = ({
  children,
  items,
  header,
}: {
  children: React.ReactNode;
  items: React.FC[];
  header?: string;
}) => {
  return (
    <div className="dropdown dropdown-end">
      <label tabIndex={0} className="btn btn-sm m-1">
        {children}
      </label>
      <ul
        tabIndex={0}
        className="dropdown-content z-[1] menu menu-sm p-2 shadow bg-base-100 rounded-box w-32"
      >
        {header && (
          <>
            <li>
              <a className="disabled text-xs font-extralight">{header}</a>
            </li>
            <li className="h-[1px] bg-gray-600/50 w-full" />
          </>
        )}
        {items.map((Item, index) => (
          <li key={index}>
            <Item />
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Dropdown;
