package main

import (
	"fmt"
	//"os"
	"strings"

	"github.com/jessefogarty/goscraper/webscraper"
)

func main() {

	/*args := os.Args[:]
	if len(args) == 2 {
		str := strings.Split(args[1], ",")
		fmt.Printf("%s", webscraper.Scraper(str))
	} else {
		os.Exit(1)
	}*/

	input := "https://www.thestar.com/news/canada/2019/08/14/ethics-report-reveals-close-co-ordination-between-trudeaus-office-and-snc-lavalin.html,https://www.thestar.com/politics/political-opinion/2019/08/14/ethics-reports-damning-findings-are-a-first-for-a-sitting-prime-minister-but-is-it-enough-to-take-trudeau-down.html,https://www.thestar.com/politics/provincial/2019/08/14/mitzie-hunter-launches-bid-for-ontario-liberal-leadership.html,https://www.thestar.com/news/gta/2019/08/14/this-toronto-homeowner-wants-to-sell-you-his-barbecue-backyard-ribs.html,https://www.thestar.com/politics/provincial/2019/08/14/ford-government-considering-overhaul-of-controversial-ontario-news-now-promotional-service.html,https://www.thestar.com/vancouver/2019/08/14/former-british-columbia-firefighter-teams-up-with-web-developer-to-create-a-wildfire-app.html"

	str := strings.Split(input, ",")

	fmt.Println(webscraper.Scraper(str))


}
