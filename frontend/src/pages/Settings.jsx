function Settings() {

  const onClick = (e) => {
    localStorage.setItem("theme", e.target.value)
  }

  return (
    <div className="flex justify-center mt-4">
      <div className="card card-md p-6 bg-base-300">
        <div className="card-title">
          <h1>Settings</h1>
        </div>
        <div className="card-body">
          <h1 className="font-bold text-lg">Change theme</h1>
          <div className="dropdown mb-72">
            <div tabIndex={0} role="button" className="btn m-1">
              Theme
              <svg
                width="12px"
                height="12px"
                className="inline-block h-2 w-2 fill-current opacity-60"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 2048 2048">
                <path d="M1799 349l242 241-1017 1017L7 590l242-241 775 775 775-775z"></path>
              </svg>
            </div>
            <ul tabIndex={0} className="dropdown-content bg-base-300 rounded-box z-1 w-52 p-2 shadow-2xl">
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  className="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Default"
                  value="dim"
                  onClick={onClick}
                />
              </li>
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  className="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Light"
                  value="light"
                  onClick={onClick}
                />
              </li>
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  className="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Retro"
                  value="retro" 
                  onClick={onClick}
                />
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Settings
