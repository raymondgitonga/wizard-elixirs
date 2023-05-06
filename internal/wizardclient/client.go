package wizardclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	GetIngredients() ([]Ingredient, error)
	GetElixirs(ingredients []string) ([]Elixir, error)
}

type Ingredient struct {
	Name string `json:"name"`
}

type Elixir struct {
	Name            string `json:"name"`
	Effect          string `json:"effect"`
	SideEffects     string `json:"sideEffects"`
	Characteristics string `json:"characteristics"`
	Difficulty      string `json:"difficulty"`
}

type RESTClient struct {
	baseURL string
	client  *http.Client
}

func NewWizardClient(baseUrl string, client *http.Client) (Client, error) {
	if client == nil {
		return nil, fmt.Errorf("http client is nil")
	}

	if len(baseUrl) < 1 {
		return nil, fmt.Errorf("wizard client url is missing")
	}

	return &RESTClient{
		client:  client,
		baseURL: baseUrl,
	}, nil
}

func (c *RESTClient) GetIngredients() ([]Ingredient, error) {
	ingredientsUrl := fmt.Sprintf("%s/ingredients", c.baseURL)

	request, err := http.NewRequest(http.MethodGet, ingredientsUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error building ingredients request: %w", err)
	}

	resp, err := c.client.Do(request)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making ingredients call: HTTP-Status %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading ingredients response: %w", err)
	}

	var ingredients []Ingredient
	err = json.Unmarshal(body, &ingredients)
	if err != nil {
		return nil, fmt.Errorf("error decoding ingredients response: %w", err)
	}
	return ingredients, nil
}

func (c *RESTClient) GetElixirs(ingredients []string) ([]Elixir, error) {
	elixirsUrl := c.buildURL(ingredients)

	request, err := http.NewRequest(http.MethodGet, elixirsUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error building elixirs request: %w", err)
	}

	resp, err := c.client.Do(request)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making elixirs call: HTTP-Status %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading elixirs response: %w", err)
	}

	var elixirs []Elixir
	err = json.Unmarshal(body, &elixirs)
	if err != nil {
		return nil, fmt.Errorf("error decoding elixirs response: %w", err)
	}
	return elixirs, nil
}

func (c *RESTClient) buildURL(ingredients []string) string {
	query := url.Values{}
	for _, v := range ingredients {
		query.Add("Ingredient", v)
	}
	return fmt.Sprintf("%s/Elixirs?%s", c.baseURL, query.Encode())
}
