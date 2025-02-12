function AvatarGroup({ avatars }) {
  return (
    <div className="avatar-group gap-1">
      {avatars.map((a) => (
        <div className="avatar avatar-placeholder">
          <div className="bg-neutral text-neutral-content w-8 rounded-full">
            <span>{a.initials}</span>
          </div>
        </div>
      ))}
    </div>
  )
}

export default AvatarGroup
