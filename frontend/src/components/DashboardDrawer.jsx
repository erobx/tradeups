import { Link } from "react-router"
import inventory from "../assets/inventory.svg"
import dashboard from "../assets/dashboard.svg"

function DashboardDrawer() {
  return (
    <div className="drawer w-min lg:drawer-open">
      <input id="my-drawer-2" type="checkbox" className="drawer-toggle" />
        <div className="drawer-content flex flex-col items-center justify-center">
          {/* Page content here */}
          <label htmlFor="my-drawer-2" className="btn btn-primary drawer-button lg:hidden">
            Open drawer
          </label>
        </div>
        <div className="drawer-side">
          <label htmlFor="my-drawer-2" aria-label="close sidebar" className="drawer-overlay"></label>
          <ul className="menu bg-base-200 text-base-content min-h-screen w-40 p-4">
            {/* Sidebar content here */}
            <li>
              <Link to="/dashboard" className="font-bold text-sm">
                <div className="flex gap-2 items-center">
                  <img src={dashboard} 
                    width={30}
                    height={30}
                    alt="Dashboard"
                  />
                  Dashboard
                </div>
              </Link>
            </li>
            <li>
              <Link to="/dashboard/inventory" className="font-bold text-sm">
                <div className="flex gap-2 items-center">
                  <img src={inventory} 
                    width={30}
                    height={30}
                    alt="Inventory"
                  />
                  Inventory
                </div>
              </Link>
            </li>
          </ul>
        </div>
    </div>
  )
}

export default DashboardDrawer
