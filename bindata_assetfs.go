package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// static_css_fonts_glyphicons_halflings_regular_eot reads file data from disk. It returns an error on failure.
func static_css_fonts_glyphicons_halflings_regular_eot() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/fonts/glyphicons-halflings-regular.eot")
	name := "static/css/fonts/glyphicons-halflings-regular.eot"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_fonts_glyphicons_halflings_regular_svg reads file data from disk. It returns an error on failure.
func static_css_fonts_glyphicons_halflings_regular_svg() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/fonts/glyphicons-halflings-regular.svg")
	name := "static/css/fonts/glyphicons-halflings-regular.svg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_fonts_glyphicons_halflings_regular_ttf reads file data from disk. It returns an error on failure.
func static_css_fonts_glyphicons_halflings_regular_ttf() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/fonts/glyphicons-halflings-regular.ttf")
	name := "static/css/fonts/glyphicons-halflings-regular.ttf"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_fonts_glyphicons_halflings_regular_woff reads file data from disk. It returns an error on failure.
func static_css_fonts_glyphicons_halflings_regular_woff() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/fonts/glyphicons-halflings-regular.woff")
	name := "static/css/fonts/glyphicons-halflings-regular.woff"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_fonts_glyphicons_halflings_regular_woff2 reads file data from disk. It returns an error on failure.
func static_css_fonts_glyphicons_halflings_regular_woff2() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/fonts/glyphicons-halflings-regular.woff2")
	name := "static/css/fonts/glyphicons-halflings-regular.woff2"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_lib_animate_css reads file data from disk. It returns an error on failure.
func static_css_lib_animate_css() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/lib/animate.css")
	name := "static/css/lib/animate.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_lib_bootstrap_theme_min_css reads file data from disk. It returns an error on failure.
func static_css_lib_bootstrap_theme_min_css() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/lib/bootstrap-theme.min.css")
	name := "static/css/lib/bootstrap-theme.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_lib_bootstrap_min_css reads file data from disk. It returns an error on failure.
func static_css_lib_bootstrap_min_css() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/lib/bootstrap.min.css")
	name := "static/css/lib/bootstrap.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_css_style_css reads file data from disk. It returns an error on failure.
func static_css_style_css() (*asset, error) {
	path := filepath.Join(rootDir, "static/css/style.css")
	name := "static/css/style.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_img_cat_png reads file data from disk. It returns an error on failure.
func static_img_cat_png() (*asset, error) {
	path := filepath.Join(rootDir, "static/img/cat.png")
	name := "static/img/cat.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_img_dog_png reads file data from disk. It returns an error on failure.
func static_img_dog_png() (*asset, error) {
	path := filepath.Join(rootDir, "static/img/dog.png")
	name := "static/img/dog.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_img_nature_png reads file data from disk. It returns an error on failure.
func static_img_nature_png() (*asset, error) {
	path := filepath.Join(rootDir, "static/img/nature.png")
	name := "static/img/nature.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_index_html reads file data from disk. It returns an error on failure.
func static_index_html() (*asset, error) {
	path := filepath.Join(rootDir, "static/index.html")
	name := "static/index.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_collections_board_js reads file data from disk. It returns an error on failure.
func static_js_app_collections_board_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/collections/board.js")
	name := "static/js/app/collections/board.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_collections_leaderboard_js reads file data from disk. It returns an error on failure.
func static_js_app_collections_leaderboard_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/collections/leaderboard.js")
	name := "static/js/app/collections/leaderboard.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_models_leaderboard_entry_js reads file data from disk. It returns an error on failure.
func static_js_app_models_leaderboard_entry_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/models/leaderboard_entry.js")
	name := "static/js/app/models/leaderboard_entry.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_models_puzzle_js reads file data from disk. It returns an error on failure.
func static_js_app_models_puzzle_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/models/puzzle.js")
	name := "static/js/app/models/puzzle.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_models_solution_js reads file data from disk. It returns an error on failure.
func static_js_app_models_solution_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/models/solution.js")
	name := "static/js/app/models/solution.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_models_tile_js reads file data from disk. It returns an error on failure.
func static_js_app_models_tile_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/models/tile.js")
	name := "static/js/app/models/tile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_templates_leaderboard_hbs reads file data from disk. It returns an error on failure.
func static_js_app_templates_leaderboard_hbs() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/templates/leaderboard.hbs")
	name := "static/js/app/templates/leaderboard.hbs"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_templates_leaderboard_entry_hbs reads file data from disk. It returns an error on failure.
func static_js_app_templates_leaderboard_entry_hbs() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/templates/leaderboard_entry.hbs")
	name := "static/js/app/templates/leaderboard_entry.hbs"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_templates_puzzle_hbs reads file data from disk. It returns an error on failure.
func static_js_app_templates_puzzle_hbs() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/templates/puzzle.hbs")
	name := "static/js/app/templates/puzzle.hbs"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_views_leaderboard_js reads file data from disk. It returns an error on failure.
func static_js_app_views_leaderboard_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/views/leaderboard.js")
	name := "static/js/app/views/leaderboard.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_views_leaderboard_entry_js reads file data from disk. It returns an error on failure.
func static_js_app_views_leaderboard_entry_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/views/leaderboard_entry.js")
	name := "static/js/app/views/leaderboard_entry.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_views_puzzle_js reads file data from disk. It returns an error on failure.
func static_js_app_views_puzzle_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/views/puzzle.js")
	name := "static/js/app/views/puzzle.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_app_views_tile_js reads file data from disk. It returns an error on failure.
func static_js_app_views_tile_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/app/views/tile.js")
	name := "static/js/app/views/tile.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_build_js reads file data from disk. It returns an error on failure.
func static_js_build_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/build.js")
	name := "static/js/build.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_lib_backbone_js reads file data from disk. It returns an error on failure.
func static_js_lib_backbone_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/lib/backbone.js")
	name := "static/js/lib/backbone.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_lib_handlebars_js reads file data from disk. It returns an error on failure.
func static_js_lib_handlebars_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/lib/handlebars.js")
	name := "static/js/lib/handlebars.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_lib_jquery_js reads file data from disk. It returns an error on failure.
func static_js_lib_jquery_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/lib/jquery.js")
	name := "static/js/lib/jquery.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_lib_require_text_js reads file data from disk. It returns an error on failure.
func static_js_lib_require_text_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/lib/require/text.js")
	name := "static/js/lib/require/text.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_lib_require_js reads file data from disk. It returns an error on failure.
func static_js_lib_require_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/lib/require.js")
	name := "static/js/lib/require.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_lib_underscore_js reads file data from disk. It returns an error on failure.
func static_js_lib_underscore_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/lib/underscore.js")
	name := "static/js/lib/underscore.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// static_js_main_js reads file data from disk. It returns an error on failure.
func static_js_main_js() (*asset, error) {
	path := filepath.Join(rootDir, "static/js/main.js")
	name := "static/js/main.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"static/css/fonts/glyphicons-halflings-regular.eot": static_css_fonts_glyphicons_halflings_regular_eot,
	"static/css/fonts/glyphicons-halflings-regular.svg": static_css_fonts_glyphicons_halflings_regular_svg,
	"static/css/fonts/glyphicons-halflings-regular.ttf": static_css_fonts_glyphicons_halflings_regular_ttf,
	"static/css/fonts/glyphicons-halflings-regular.woff": static_css_fonts_glyphicons_halflings_regular_woff,
	"static/css/fonts/glyphicons-halflings-regular.woff2": static_css_fonts_glyphicons_halflings_regular_woff2,
	"static/css/lib/animate.css": static_css_lib_animate_css,
	"static/css/lib/bootstrap-theme.min.css": static_css_lib_bootstrap_theme_min_css,
	"static/css/lib/bootstrap.min.css": static_css_lib_bootstrap_min_css,
	"static/css/style.css": static_css_style_css,
	"static/img/cat.png": static_img_cat_png,
	"static/img/dog.png": static_img_dog_png,
	"static/img/nature.png": static_img_nature_png,
	"static/index.html": static_index_html,
	"static/js/app/collections/board.js": static_js_app_collections_board_js,
	"static/js/app/collections/leaderboard.js": static_js_app_collections_leaderboard_js,
	"static/js/app/models/leaderboard_entry.js": static_js_app_models_leaderboard_entry_js,
	"static/js/app/models/puzzle.js": static_js_app_models_puzzle_js,
	"static/js/app/models/solution.js": static_js_app_models_solution_js,
	"static/js/app/models/tile.js": static_js_app_models_tile_js,
	"static/js/app/templates/leaderboard.hbs": static_js_app_templates_leaderboard_hbs,
	"static/js/app/templates/leaderboard_entry.hbs": static_js_app_templates_leaderboard_entry_hbs,
	"static/js/app/templates/puzzle.hbs": static_js_app_templates_puzzle_hbs,
	"static/js/app/views/leaderboard.js": static_js_app_views_leaderboard_js,
	"static/js/app/views/leaderboard_entry.js": static_js_app_views_leaderboard_entry_js,
	"static/js/app/views/puzzle.js": static_js_app_views_puzzle_js,
	"static/js/app/views/tile.js": static_js_app_views_tile_js,
	"static/js/build.js": static_js_build_js,
	"static/js/lib/backbone.js": static_js_lib_backbone_js,
	"static/js/lib/handlebars.js": static_js_lib_handlebars_js,
	"static/js/lib/jquery.js": static_js_lib_jquery_js,
	"static/js/lib/require/text.js": static_js_lib_require_text_js,
	"static/js/lib/require.js": static_js_lib_require_js,
	"static/js/lib/underscore.js": static_js_lib_underscore_js,
	"static/js/main.js": static_js_main_js,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"css": &_bintree_t{nil, map[string]*_bintree_t{
			"fonts": &_bintree_t{nil, map[string]*_bintree_t{
				"glyphicons-halflings-regular.eot": &_bintree_t{static_css_fonts_glyphicons_halflings_regular_eot, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.svg": &_bintree_t{static_css_fonts_glyphicons_halflings_regular_svg, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.ttf": &_bintree_t{static_css_fonts_glyphicons_halflings_regular_ttf, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.woff": &_bintree_t{static_css_fonts_glyphicons_halflings_regular_woff, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.woff2": &_bintree_t{static_css_fonts_glyphicons_halflings_regular_woff2, map[string]*_bintree_t{
				}},
			}},
			"lib": &_bintree_t{nil, map[string]*_bintree_t{
				"animate.css": &_bintree_t{static_css_lib_animate_css, map[string]*_bintree_t{
				}},
				"bootstrap-theme.min.css": &_bintree_t{static_css_lib_bootstrap_theme_min_css, map[string]*_bintree_t{
				}},
				"bootstrap.min.css": &_bintree_t{static_css_lib_bootstrap_min_css, map[string]*_bintree_t{
				}},
			}},
			"style.css": &_bintree_t{static_css_style_css, map[string]*_bintree_t{
			}},
		}},
		"img": &_bintree_t{nil, map[string]*_bintree_t{
			"cat.png": &_bintree_t{static_img_cat_png, map[string]*_bintree_t{
			}},
			"dog.png": &_bintree_t{static_img_dog_png, map[string]*_bintree_t{
			}},
			"nature.png": &_bintree_t{static_img_nature_png, map[string]*_bintree_t{
			}},
		}},
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			"app": &_bintree_t{nil, map[string]*_bintree_t{
				"collections": &_bintree_t{nil, map[string]*_bintree_t{
					"board.js": &_bintree_t{static_js_app_collections_board_js, map[string]*_bintree_t{
					}},
					"leaderboard.js": &_bintree_t{static_js_app_collections_leaderboard_js, map[string]*_bintree_t{
					}},
				}},
				"models": &_bintree_t{nil, map[string]*_bintree_t{
					"leaderboard_entry.js": &_bintree_t{static_js_app_models_leaderboard_entry_js, map[string]*_bintree_t{
					}},
					"puzzle.js": &_bintree_t{static_js_app_models_puzzle_js, map[string]*_bintree_t{
					}},
					"solution.js": &_bintree_t{static_js_app_models_solution_js, map[string]*_bintree_t{
					}},
					"tile.js": &_bintree_t{static_js_app_models_tile_js, map[string]*_bintree_t{
					}},
				}},
				"templates": &_bintree_t{nil, map[string]*_bintree_t{
					"leaderboard.hbs": &_bintree_t{static_js_app_templates_leaderboard_hbs, map[string]*_bintree_t{
					}},
					"leaderboard_entry.hbs": &_bintree_t{static_js_app_templates_leaderboard_entry_hbs, map[string]*_bintree_t{
					}},
					"puzzle.hbs": &_bintree_t{static_js_app_templates_puzzle_hbs, map[string]*_bintree_t{
					}},
				}},
				"views": &_bintree_t{nil, map[string]*_bintree_t{
					"leaderboard.js": &_bintree_t{static_js_app_views_leaderboard_js, map[string]*_bintree_t{
					}},
					"leaderboard_entry.js": &_bintree_t{static_js_app_views_leaderboard_entry_js, map[string]*_bintree_t{
					}},
					"puzzle.js": &_bintree_t{static_js_app_views_puzzle_js, map[string]*_bintree_t{
					}},
					"tile.js": &_bintree_t{static_js_app_views_tile_js, map[string]*_bintree_t{
					}},
				}},
			}},
			"build.js": &_bintree_t{static_js_build_js, map[string]*_bintree_t{
			}},
			"lib": &_bintree_t{nil, map[string]*_bintree_t{
				"backbone.js": &_bintree_t{static_js_lib_backbone_js, map[string]*_bintree_t{
				}},
				"handlebars.js": &_bintree_t{static_js_lib_handlebars_js, map[string]*_bintree_t{
				}},
				"jquery.js": &_bintree_t{static_js_lib_jquery_js, map[string]*_bintree_t{
				}},
				"require": &_bintree_t{nil, map[string]*_bintree_t{
					"text.js": &_bintree_t{static_js_lib_require_text_js, map[string]*_bintree_t{
					}},
				}},
				"require.js": &_bintree_t{static_js_lib_require_js, map[string]*_bintree_t{
				}},
				"underscore.js": &_bintree_t{static_js_lib_underscore_js, map[string]*_bintree_t{
				}},
			}},
			"main.js": &_bintree_t{static_js_main_js, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: k}
	}
	panic("unreachable")
}
