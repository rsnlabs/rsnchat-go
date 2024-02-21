package rsnchatgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Result represents the generic response structure from the API.
type Result struct {
	Success string `json:"success"`
}

// TextResult represents the response structure for text-based API endpoints.
type TextResult struct {
	Result
	Message string `json:"message"`
}

// ProdiaResult represents the response structure for Prodia API endpoint.
type ProdiaResult struct {
	Result
	ImageURL string `json:"imageUrl"`
	Base64   string `json:"base64"`
}

// Image represents the response structure for image-based API endpoints.
type Image struct {
	Result
	Image string `json:"image"`
}

// ProdiaModels represents the available Prodia models.
type ProdiaModels string

// Constants for Prodia models.
const (
	Model3Guofeng3V34                                 ProdiaModels = "3Guofeng3_v34.safetensors [50f420de]"
	ModelAbsoluterealityV16                           ProdiaModels = "absolutereality_V16.safetensors [37db0fc3]"
	ModelAbsoluterealityV181                          ProdiaModels = "absolutereality_v181.safetensors [3d9d4d2b]"
	ModelAmIRealV41                                   ProdiaModels = "amIReal_V41.safetensors [0a8a2e61]"
	ModelAnalogDiffusion10Ckpt                        ProdiaModels = "analog-diffusion-1.0.ckpt [9ca13f02]"
	ModelAnythingv300PrunedCkpt                       ProdiaModels = "anythingv3_0-pruned.ckpt [2700c435]"
	ModelAnythingV45PrunedCkpt                        ProdiaModels = "anything-v4.5-pruned.ckpt [65745d25]"
	ModelAnythingV5PrtRESafetensors                   ProdiaModels = "anythingV5_PrtRE.safetensors [893e49b9]"
	ModelAOM3A3OrangemixsSafetensors                  ProdiaModels = "AOM3A3_orangemixs.safetensors [9600da17]"
	ModelBlazingDriveV10gSafetensors                  ProdiaModels = "blazing_drive_v10g.safetensors [ca1c1eab]"
	ModelCetusMixVersion35Safetensors                 ProdiaModels = "cetusMix_Version35.safetensors [de2f2560]"
	ModelChildrensStoriesV13DSafetensors              ProdiaModels = "childrensStories_v13D.safetensors [9dfaabcb]"
	ModelChildrensStoriesV1SemiRealSafetensors        ProdiaModels = "childrensStories_v1SemiReal.safetensors [a1c56dbb]"
	ModelChildrensStoriesV1ToonAnimeSafetensors       ProdiaModels = "childrensStories_v1ToonAnime.safetensors [2ec7b88b]"
	ModelCounterfeitV30Safetensors                    ProdiaModels = "Counterfeit_v30.safetensors [9e2a8f19]"
	ModelCuteyukimixAdorableMidchapter3Safetensors    ProdiaModels = "cuteyukimixAdorable_midchapter3.safetensors [04bdffe6]"
	ModelCyberrealisticV33Safetensors                 ProdiaModels = "cyberrealistic_v33.safetensors [82b0d085]"
	ModelDalcefoV4Safetensors                         ProdiaModels = "dalcefo_v4.safetensors [425952fe]"
	ModelDeliberateV2Safetensors                      ProdiaModels = "deliberate_v2.safetensors [10ec4b29]"
	ModelDeliberateV3Safetensors                      ProdiaModels = "deliberate_v3.safetensors [afd9d2d4]"
	ModelDreamlikeAnime10Safetensors                  ProdiaModels = "dreamlike-anime-1.0.safetensors [4520e090]"
	ModelDreamlikeDiffusion10Safetensors              ProdiaModels = "dreamlike-diffusion-1.0.safetensors [5c9fd6e0]"
	ModelDreamlikePhotoreal20Safetensors              ProdiaModels = "dreamlike-photoreal-2.0.safetensors [fdcf65e7]"
	ModelDreamshaper6BakedVaeSafetensors              ProdiaModels = "dreamshaper_6BakedVae.safetensors [114c8abb]"
	ModelDreamshaper7Safetensors                      ProdiaModels = "dreamshaper_7.safetensors [5cf5ae06]"
	ModelDreamshaper8Safetensors                      ProdiaModels = "dreamshaper_8.safetensors [9d40847d]"
	ModelEdgeOfRealismEorV20Safetensors               ProdiaModels = "edgeOfRealism_eorV20.safetensors [3ed5de15]"
	ModelEimisAnimeDiffusionV1Ckpt                    ProdiaModels = "EimisAnimeDiffusion_V1.ckpt [4f828a15]"
	ModelElldrethsVividMixSafetensors                 ProdiaModels = "elldreths-vivid-mix.safetensors [342d9d26]"
	ModelEpicrealismNaturalSinRC1VaESafetensors       ProdiaModels = "epicrealism_naturalSinRC1VAE.safetensors [90a4c676]"
	ModelICantBelieveItsNotPhotographySecoSafetensors ProdiaModels = "ICantBelieveItsNotPhotography_seco.safetensors [4e7a3dfd]"
	ModelJuggernautAftermathSafetensors               ProdiaModels = "juggernaut_aftermath.safetensors [5e20c455]"
	ModelLofiV4Safetensors                            ProdiaModels = "lofi_v4.safetensors [ccc204d6]"
	ModelLyrielV16Safetensors                         ProdiaModels = "lyriel_v16.safetensors [68fceea2]"
	ModelMajicmixRealisticV4Safetensors               ProdiaModels = "majicmixRealistic_v4.safetensors [29d0de58]"
	ModelMechamixV10Safetensors                       ProdiaModels = "mechamix_v10.safetensors [ee685731]"
	ModelMeinamixMeinaV9Safetensors                   ProdiaModels = "meinamix_meinaV9.safetensors [2ec66ab0]"
	ModelMeinamixMeinaV11Safetensors                  ProdiaModels = "meinamix_meinaV11.safetensors [b56ce717]"
	ModelNeverendingDreamV122Safetensors              ProdiaModels = "neverendingDream_v122.safetensors [f964ceeb]"
	ModelOpenjourneyV4Ckpt                            ProdiaModels = "openjourney_V4.ckpt [ca2f377f]"
	ModelPastelMixStylizedAnimePrunedFp16Safetensors  ProdiaModels = "pastelMixStylizedAnime_pruned_fp16.safetensors [793a26e8]"
	ModelPortraitplusV10Safetensors                   ProdiaModels = "portraitplus_V1.0.safetensors [1400e684]"
	ModelProtogenx34Safetensors                       ProdiaModels = "protogenx34.safetensors [5896f8d5]"
	ModelRealisticVisionV14PrunedFp16Safetensors      ProdiaModels = "Realistic_Vision_V1.4-pruned-fp16.safetensors [8d21810b]"
	ModelRealisticVisionV20Safetensors                ProdiaModels = "Realistic_Vision_V2.0.safetensors [79587710]"
	ModelRealisticVisionV40Safetensors                ProdiaModels = "Realistic_Vision_V4.0.safetensors [29a7afaa]"
	ModelRealisticVisionV50Safetensors                ProdiaModels = "Realistic_Vision_V5.0.safetensors [614d1063]"
	ModelRedshiftDiffusionV10Safetensors              ProdiaModels = "redshift_diffusion-V10.safetensors [1400e684]"
	ModelRevAnimatedV122Safetensors                   ProdiaModels = "revAnimated_v122.safetensors [3f4fefd9]"
	ModelRundiffusionFX25DV10Safetensors              ProdiaModels = "rundiffusionFX25D_v10.safetensors [cd12b0ee]"
	ModelRundiffusionFXV10Safetensors                 ProdiaModels = "rundiffusionFX_v10.safetensors [cd4e694d]"
	ModelSdv14Ckpt                                    ProdiaModels = "sdv1_4.ckpt [7460a6fa]"
	ModelShoninsBeautifulV10Safetensors               ProdiaModels = "shoninsBeautiful_v10.safetensors [25d8c546]"
	ModelTheallysMixIiChurnedSafetensors              ProdiaModels = "theallys-mix-ii-churned.safetensors [5d9225a4]"
	ModelTimeless10Ckpt                               ProdiaModels = "timeless-1.0.ckpt [7c4971d4]"
	ModelToonyouBeta6Safetensors                      ProdiaModels = "toonyou_beta6.safetensors [980f6b15]"
)

// RsnChat represents the RsnChat client.
type RsnChat struct {
	APIKey  string
	APIURL  string
	Headers map[string]string
	Client  *http.Client
}

// NewRsnChat creates a new RsnChat client with the provided API key.
func NewRsnChat(apiKey string, apiURL ...string) (*RsnChat, error) {
	var trueApiURL string

	if len(apiURL) == 0 {
		trueApiURL = "https://api.rsnai.org/api/v1/user"
	} else if len(apiURL) == 1 {
		trueApiURL = apiURL[0]
	} else {
		return nil, errors.New("too many arguments")
	}

	if apiKey == "" {
		return nil, errors.New("please provide API key")
	}

	// Validate API Key
	validateURL := fmt.Sprintf("%s/validate", trueApiURL)
	reqBody, err := json.Marshal(map[string]string{"key": apiKey})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(validateURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Check if the API key is valid
	body, err := io.ReadAll(resp.Body)
	if err != nil || !strings.Contains(string(body), "API key validated") {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid API Key: %s", apiKey)
	}

	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", apiKey)}
	return &RsnChat{APIKey: apiKey, APIURL: trueApiURL, Headers: headers, Client: http.DefaultClient}, nil
}

// gpt sends a request to the GPT API endpoint.
func (r *RsnChat) Gpt(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/gpt", payload)
}

// openchat sends a request to the OpenChat API endpoint.
func (r *RsnChat) Openchat(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/openchat", payload)
}

// bard sends a request to the Bard API endpoint.
func (r *RsnChat) Bard(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/bard", payload)
}

// gemini sends a request to the Gemini API endpoint.
func (r *RsnChat) Gemini(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/gemini", payload)
}

// bing sends a request to the Bing API endpoint.
func (r *RsnChat) Bing(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/bing", payload)
}

// llama sends a request to the Llama API endpoint.
func (r *RsnChat) Llama(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/llama", payload)
}

// mixtral sends a request to the Mixtral API endpoint.
func (r *RsnChat) Mixtral(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/mixtral", payload)
}

// claude sends a request to the Claude API endpoint.
func (r *RsnChat) Claude(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/claude", payload)
}

// codellama sends a request to the Codellama API endpoint.
func (r *RsnChat) Codellama(prompt string) (*TextResult, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendRequest("/codellama", payload)
}

// prodia sends a request to the Prodia API endpoint.
func (r *RsnChat) Prodia(prompt string, negativePrompt string, model string) (*Image, error) {
	payload := map[string]string{"prompt": prompt, "negative_prompt": negativePrompt, "model": model}
	return r.sendImageRequest("/prodia", payload)
}

// kandinsky sends a request to the Kandinsky API endpoint.
func (r *RsnChat) Kandinsky(prompt string, negativePrompt string) (*Image, error) {
	payload := map[string]string{"prompt": prompt, "negative_prompt": negativePrompt}
	return r.sendImageRequest("/kandinsky", payload)
}

// absolutebeauty sends a request to the AbsoluteBeauty API endpoint.
func (r *RsnChat) Absolutebeauty(prompt string, negativePrompt string) (*Image, error) {
	payload := map[string]string{"prompt": prompt, "negative_prompt": negativePrompt}
	return r.sendImageRequest("/absolutebeauty", payload)
}

func (r *RsnChat) Sdxl(prompt string, negativePrompt string) (*Image, error) {
	payload := map[string]string{"prompt": prompt, "negative_prompt": negativePrompt}
	return r.sendImageRequest("/sdxl", payload)
}

func (r *RsnChat) Dalle(prompt string) (*Image, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendImageRequest("/dalle", payload)
}

func (r *RsnChat) Icon(prompt string) (*Image, error) {
	payload := map[string]string{"prompt": prompt}
	return r.sendImageRequest("/icon", payload)
}

func (r *RsnChat) sendRequest(endpoint string, payload map[string]string) (*TextResult, error) {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Post(r.APIURL+endpoint, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handling specific HTTP status codes
	switch resp.StatusCode {
	case http.StatusOK:
		var result TextResult
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: API key is invalid or you do not have access to the resource")
	case http.StatusForbidden:
		return nil, errors.New("forbidden: access to the resource is denied")
	case http.StatusNotFound:
		return nil, errors.New("not found: the requested endpoint is not available")
	default:
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}
}

func (r *RsnChat) sendImageRequest(endpoint string, payload map[string]string) (*Image, error) {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", r.APIURL+endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.APIKey)

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handling specific HTTP status codes
	switch resp.StatusCode {
	case http.StatusOK:
		var result Image
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: API key is invalid")
	case http.StatusForbidden:
		return nil, errors.New("forbidden: access to the resource is denied")
	case http.StatusNotFound:
		return nil, errors.New("not found: the requested endpoint is not available")
	default:
		return nil, fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}
}
