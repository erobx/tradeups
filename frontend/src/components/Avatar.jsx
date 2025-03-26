import useUser from "../stores/userStore"

function Avatar() {
  const { user, setUser, setBalance } = useUser()

  return (
    <div tabIndex={0} role="button" className="btn btn-ghost btn-circle ml-1 avatar avatar-placeholder">
      <div className="bg-neutral text-neutral-content w-12 rounded-full">
        <span>{user.username[0].toUpperCase()}</span>
      </div>
    </div>
  )
}

export default Avatar
