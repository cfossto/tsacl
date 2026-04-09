package tailscale

const tailnetBaseUrl = "https://api.tailscale.com/api/v2/tailnet/"

func CurrentTailnetContext(tailnet string) string {
	return tailnetBaseUrl + tailnet + "/acl"
}


