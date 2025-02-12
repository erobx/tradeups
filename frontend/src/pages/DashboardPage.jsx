import { Routes, Route } from "react-router"

import DashboardDrawer from "../components/DashboardDrawer"
import Inventory from "../components/Inventory"
import Dashboard from "../components/Dashboard"

function DashboardPage() {
  return (
    <div className="flex gap-5">
      <div>
        <DashboardDrawer />
      </div>
      <div className="mt-4">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/inventory" element={<Inventory />} />
        </Routes>
      </div>
    </div>
  )
}

export default DashboardPage
