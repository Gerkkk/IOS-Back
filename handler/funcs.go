package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Gerkkk/IOS-Back/entities"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	_ "mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

var users []entities.Person = []entities.Person{
	entities.Person{
		Id:           0,
		Name:         "Patrick Bateman",
		Tag:          "patrickstar",
		Status:       "Sigma Sigma boy sigma boy sigma boy",
		Follows:      []int{1, 2},
		FollowersNum: 777,
		Avatar:       "/avatars/patrickbateman.jpeg",
		Likes:        []int{2, 3, 4},
	}, entities.Person{
		Id:           1,
		Name:         "Boba Bibov",
		Tag:          "biba",
		Status:       "Living life",
		Follows:      []int{0, 2},
		FollowersNum: 66,
		Avatar:       "",
	}, entities.Person{
		Id:           2,
		Name:         "Chris",
		Tag:          "boba",
		Status:       "Making love",
		Follows:      []int{0, 1},
		FollowersNum: 0,
		Avatar:       "",
	}, entities.Person{
		Id:           3,
		Name:         "Серж Дур-Дачник",
		Tag:          "riverwalker",
		Status:       "Я в своем познании настолько преисполнился",
		Follows:      []int{0, 1},
		FollowersNum: 1000000000,
		Avatar:       "/avatars/riverwalker.jpg",
	},
}

var articles []entities.Article = []entities.Article{
	entities.Article{
		ID:           0,
		AuthorID:     0,
		Title:        "Trip to New York",
		AuthorName:   "Patrick Bateman",
		AuthorTag:    "patrickstar",
		AuthorAvatar: "/avatars/patrickbateman.jpeg",
		Date:         time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local).Unix(),
		CoordsN:      40.730610,
		CoordsW:      -73.935242,
		Brief:        "Had lots of fun there",
		Text:         "New York is a jungle, a shimmering facade of wealth and status. The streets pulse with ambition, suffocating in its manic energy. The noise is maddening, but I thrive in it. The people are nothing more than faceless commodities—empty, hollow, just like me. The city’s soul is mine to devour.",
		Images:       []string{"/images/newyork1.avif", "/images/newyork2.webp"},
	},
	entities.Article{
		ID:           1,
		AuthorID:     0,
		AuthorName:   "Patrick Bateman",
		AuthorTag:    "patrickstar",
		AuthorAvatar: "/avatars/patrickbateman.jpeg",
		Title:        "Trip to Chicago",
		Date:         time.Date(2010, time.January, 10, 23, 0, 0, 0, time.UTC).Unix(),
		CoordsN:      42.0,
		CoordsW:      -87.6,
		Brief:        "Raw city",
		Text:         "Chicago is raw, untamed—an industrial beast beneath the gleaming skyline. The wind howls through the streets, cold and unforgiving, like the city itself. People here are tougher, more direct. There's a certain grit to it, a brutal honesty that feels almost… refreshing. A place where only the strong survive.",
		Images:       []string{"/images/chicago1.jpg", "/images/chicago2.jpeg"},
	},
	entities.Article{
		ID:           2,
		AuthorID:     3,
		AuthorName:   "Серж Дур-Дачник",
		AuthorTag:    "riverwalker",
		AuthorAvatar: "/avatars/riverwalker.jpg",
		Title:        "Прогулка к реке",
		Date:         time.Date(2015, time.January, 10, 23, 0, 0, 0, time.UTC).Unix(),
		CoordsN:      45.0,
		CoordsW:      38.5,
		Brief:        "Вспыхнуло солнце закатом\nНа дворе благодать\nМир абсолютно понятен, но тебе не понять\nНебо целуют кометы, ты смеешься опять\nМир абсолютно понятен, но тебе не понять",
		Text:         "Я в своем познании настолько преисполнился, что я как будто бы уже сто триллионов миллиардов лет проживаю на триллионах и триллионах таких же планет, как эта Земля, мне этот мир абсолютно понятен, и я здесь ищу только одного - покоя, умиротворения и вот этой гармонии, от слияния с бесконечно вечным, от созерцания великого фрактального подобия и от вот этого замечательного всеединства существа, бесконечно вечного, куда ни посмотри, хоть вглубь - бесконечно малое, хоть ввысь - бесконечное большое, понимаешь? А ты мне опять со своим вот этим, иди суетись дальше, это твоё распределение, это твой путь и твой горизонт познания и ощущения твоей природы, он несоизмеримо мелок по сравнению с моим, понимаешь? Я как будто бы уже давно глубокий старец, бессмертный, ну или там уже почти бессмертный, который на этой планете от её самого зарождения, ещё когда только Солнце только-только сформировалось как звезда, и вот это газопылевое облако, вот, после взрыва, Солнца, когда оно вспыхнуло, как звезда, начало формировать вот эти коацерваты, планеты, понимаешь, я на этой Земле уже как будто почти пять миллиардов лет живу и знаю её вдоль и поперёк этот весь мир, а ты мне какие-то... мне не важно на твои тачки, на твои яхты, на твои квартиры, там, на твоё благо. Я был на этой планете бесконечным множеством, и круче Цезаря, и круче Гитлера, и круче всех великих, понимаешь, был, а где-то был конченым говном, ещё хуже, чем здесь. Я множество этих состояний чувствую. Где-то я был больше подобен растению, где-то я больше был подобен птице, там, червю, где-то был просто сгусток камня, это всё есть душа, понимаешь? Она имеет грани подобия совершенно многообразные, бесконечное множество. Но тебе этого не понять, поэтому ты езжай себе , мы в этом мире как бы живем разными ощущениями и разными стремлениями, соответственно, разное наше и место, разное и наше распределение. Тебе я желаю все самые крутые тачки чтоб были у тебя, и все самые лучше самки, если мало идей, обращайся ко мне, я тебе на каждую твою идею предложу сотню триллионов, как всё делать. Ну а я всё, я иду как глубокий старец,узревший вечное, прикоснувшийся к Божественному, сам стал богоподобен и устремлен в это бесконечное, и который в умиротворении, покое, гармонии, благодати, в этом сокровенном блаженстве пребывает, вовлеченный во всё и во вся, понимаешь, вот и всё, в этом наша разница. Так что я иду любоваться мирозданием, а ты идёшь преисполняться в ГРАНЯХ каких-то, вот и вся разница, понимаешь, ты не зришь это вечное бесконечное, оно тебе не нужно. Ну зато ты, так сказать, более активен, как вот этот дятел долбящий, или муравей, который очень активен в своей стезе, поэтому давай, наши пути здесь, конечно, имеют грани подобия, потому что всё едино, но я-то тебя прекрасно понимаю, а вот ты меня - вряд ли, потому что я как бы тебя в себе содержу, всю твою природу, она составляет одну маленькую там песчиночку, от того что есть во мне, вот и всё, поэтому давай, ступай, езжай, а я пошел наслаждаться прекрасным осенним закатом на берегу теплой южной реки. Всё, ступай, и я пойду.",
		Images:       []string{"/images/riverwalker1.jpg"},
	},
	entities.Article{
		ID:         3,
		AuthorID:   2,
		AuthorName: "Chris",
		AuthorTag:  "boba",
		Title:      "Love Valley",
		Date:       time.Date(2020, time.July, 10, 10, 0, 0, 0, time.UTC).Unix(),
		CoordsN:    38.6,
		CoordsW:    35.0,
		Brief:      "Wonderful journey to Cappadocia",
		Text:       "The history of Love Valley dates back to at least Roman times. There goes a legend that there once was two dynasties living in the same village. A fight broke out between the two dynasties, which resulted in the village effectively being split. One day, two villagers complained about the situation which resulted in the recruitment of two people from opposing sides. The two recruited soldiers fell in love with each other as soon as they saw each other. The feuding villagers, having had knowledge of this, tried their best to separate the two but failed. After they struggled to separate the two, the villagers decided to get them married. Time passed, the couple had a child, however the situation wasn't enough to reconcile the opposing families. Finally, they killed the boy. The girl couldn't stand her husband's death and later committed suicide. It is said that after the death of the two lovers, God rained stones to punish the feuding villagers. These stones are to kill anyone who opposes the reunion of youth.",
		Images:     []string{"/images/lovevalley1.jpg", "/images/lovevalley2.jpg"},
	},
	entities.Article{
		ID:         4,
		AuthorID:   1,
		AuthorName: "Boba Bibov",
		AuthorTag:  "biba",
		Title:      "No place like London",
		Date:       time.Date(2020, time.July, 10, 10, 0, 0, 0, time.UTC).Unix(),
		CoordsN:    51.509865,
		CoordsW:    -0.118092,
		Brief:      "Foggy capital city",
		Text:       "London is a city of contrasts, where history meets modernity. Its streets pulse with energy, from the iconic red buses to the quiet corners of ancient pubs. The skyline is a mix of glass towers and centuries-old landmarks. It's a place that constantly reinvents itself, yet never forgets its roots.",
		Images:     []string{"/images/london1.webp", "/images/london2.jpg"},
	},
}

func (h *Handler) search(c *gin.Context) {
	input := new(entities.SearchStruct)
	err := c.BindJSON(&input)

	if err != nil {
		log.Fatalf("error1")
	}

	if input.SearchingType == 0 {

		c.JSON(http.StatusOK, map[string]interface{}{
			"people": users,
		})
	} else {

		c.JSON(http.StatusOK, map[string]interface{}{
			"articles": articles,
		})
	}
}

func (h *Handler) getNews(c *gin.Context) {
	var input entities.NewsStruct
	_ = c.BindJSON(&input)

	fmt.Println("! ", input.UserId)

	if (input.PageNum-1)*input.PageSize > len(articles)-1 {
		ret := []entities.Article{}
		c.JSON(http.StatusOK, map[string]interface{}{
			"articles": ret,
		})
		return
	}

	ret := articles[(input.PageNum-1)*input.PageSize : min(input.PageNum*input.PageSize, len(articles))]

	c.JSON(http.StatusOK, map[string]interface{}{
		"articles": ret,
	})
}

func (h *Handler) likePost(c *gin.Context) {
	var input entities.LikeStruct
	_ = c.BindJSON(&input)

	if input.UserID < len(users) {
		users[input.UserID].Likes = append(users[input.UserID].Likes, input.PostID)
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) follow(c *gin.Context) {
	var input entities.LikeStruct
	_ = c.BindJSON(input)

	if input.UserID < len(users) {
		users[input.UserID].Follows = append(users[input.UserID].Follows, input.PostID)
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) getLikedPosts(c *gin.Context) {
	var input entities.PersonIDStruct

	_ = c.BindJSON(&input)

	if input.ID < len(users) {
		var ret []entities.Article
		for _, i := range users[input.ID].Likes {
			ret = append(ret, articles[i])
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"articles": ret,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
	}
}

func (h *Handler) getUserPosts(c *gin.Context) {
	var input entities.PersonIDStruct

	_ = c.BindJSON(&input)

	if input.ID < len(users) {
		var ret []entities.Article
		for _, i := range articles {
			if i.AuthorID == input.ID {
				ret = append(ret, i)
			}
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"articles": ret,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
	}
}

func (h *Handler) getUserInfo(c *gin.Context) {
	var input entities.PersonIDStruct
	_ = c.BindJSON(&input)

	if input.ID < len(users) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"user": users[input.ID],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
	}
}

func (h *Handler) register(c *gin.Context) {
	var input entities.RegisterPerson

	_ = c.BindJSON(&input)

	var newPerson entities.Person

	newPerson.Name = input.Name
	newPerson.Tag = input.Tag

	users = append(users, newPerson)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":           len(users) - 1,
		"accessToken":  "lalolilelo",
		"refreshToken": "lalolilelo",
	})
}

func (h *Handler) login(c *gin.Context) {
	var input entities.LoginPerson
	_ = c.BindJSON(&input)

	fmt.Println(input)

	var ind = -1

	for i, person := range users {
		if person.Tag == input.Tag {
			ind = i
			break
		}
	}

	if ind != -1 {
		fmt.Println(ind)
		c.JSON(http.StatusOK, map[string]interface{}{
			"id":           ind,
			"accessToken":  "lalolilelo",
			"refreshToken": "lalolilelo",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
	}
}

func (h *Handler) refresh(c *gin.Context) {
	var input entities.RefreshStruct
	_ = c.BindJSON(&input)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  "lalolilelo",
		"refreshToken": "lalolilelo",
	})
}

func (h *Handler) createNewArticle(c *gin.Context) {
	var articleData entities.Article

	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse multipart form"})
		return
	}

	//jsonFile := form.File["json"]
	jsonFile, err := c.FormFile("json")
	jsonData, _ := jsonFile.Open()
	defer jsonData.Close()
	fileBytes, err := io.ReadAll(jsonData)

	if err := json.Unmarshal([]byte(fileBytes), &articleData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	articleData.ID = len(articles)

	if articleData.AuthorID < len(users) {
		articleData.AuthorName = users[articleData.AuthorID].Name
		articleData.AuthorTag = users[articleData.AuthorID].Tag
		articleData.AuthorAvatar = users[articleData.AuthorID].Avatar
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
	}

	files := form.File["images"]
	for _, file := range files {
		filePath := filepath.Join("./images", file.Filename)
		articleData.Images = append(articleData.Images, filepath.Join("/images", file.Filename))
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
			return
		}
	}

	articles = append(articles, articleData)

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) changeSettings(c *gin.Context) {
	var userData entities.Person

	jsonFile, _ := c.FormFile("json")
	jsonData, _ := jsonFile.Open()
	defer jsonData.Close()
	fileBytes, _ := io.ReadAll(jsonData)

	if err := json.Unmarshal(fileBytes, &userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if userData.Id < len(users) {
		users[userData.Id].Tag = userData.Tag
		users[userData.Id].Status = userData.Status
		users[userData.Id].Name = userData.Name
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
	}

	file, _ := c.FormFile("avatar")

	filePath := filepath.Join("./avatars", file.Filename)
	users[userData.Id].Avatar = filepath.Join("/avatars", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
		return
	}

	for i := range articles {
		var article = &articles[i]
		if article.AuthorID == userData.Id {
			article.AuthorName = userData.Name
			article.AuthorTag = userData.Tag
			article.AuthorAvatar = users[userData.Id].Avatar
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) getSettings(c *gin.Context) {
	var input entities.PersonIDStruct
	_ = c.BindJSON(&input)
	if input.ID < len(users) {
		var ret = entities.Settings{}
		ret.Id = users[input.ID].Id
		ret.Name = users[input.ID].Name
		ret.Tag = users[input.ID].Tag
		ret.Status = users[input.ID].Status
		ret.Avatar = users[input.ID].Avatar

		c.JSON(http.StatusOK, map[string]interface{}{
			"user": ret,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{})
	}
}
