package convert

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/dietsche/rfsnotify"
	"github.com/rs/zerolog/log"
)

func StartWatcher(sourceDir, destDir string) error {
	// watcher, err := rfsnotify.NewWatcher()
	watcher, err := rfsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()
	watcher.AddRecursive(sourceDir)
	// 使用 WaitGroup 以确保所有子 goroutine 完成后再退出程序
	var wg sync.WaitGroup
	wg.Add(1)
	// 启动 goroutine 监听事件
	go func() {
		defer wg.Done()
		for {
			select {
			case event := <-watcher.Events:
				if filepath.Ext(event.Name) == ".md" {
					// .md 文件被写入，进行同步操作
					relPath, err := filepath.Rel(sourceDir, event.Name)
					if err != nil {
						continue
					}
					destPath := filepath.Join(destDir, relPath[:len(relPath)-len(".md")]+".wiki")
					// 确保目标目录及其上级目录存在
					err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
					if err != nil {
						continue
					}
					ConvertMd2Json(event.Name, destPath)
					RenderLayout(sourceDir, destDir)
				}

			case err := <-watcher.Errors:
				log.Error().Err(err).Msg("moniter watch error")

			}
		}
	}()

	// 添加监听源目录及其子目录
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				watcher.Add(path)
			}()
		}
		return nil
	})
	if err != nil {
		return err
	}
	log.Info().Msg("Watcher started. Press Ctrl+C to stop.")
	// 等待所有子 goroutine 完成
	wg.Wait()
	return nil
}
