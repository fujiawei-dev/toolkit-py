---
date: {{ created_at }}  # 创建日期
author: "Rustle Karl"  # 作者

title: "{{ project_slug.words_capitalized }} 学习笔记"  # 文章标题
description: "纸上得来终觉浅，学到过知识点分分钟忘得一干二净，今后无论学什么，都做好笔记吧。"
url:  "posts/{{ project_slug.kebab_case }}/README"  # 设置网页永久链接
tags: [ "{{ project_slug.kebab_case }}", "README" ]  # 标签
categories: [ "{{ project_slug.words_capitalized }} 学习笔记" ]  # 分类

index: true  # 是否可以被索引
toc: true  # 是否自动生成目录
draft: false  # 草稿
---

# {{ project_slug.words_capitalized }} 学习笔记

> 纸上得来终觉浅，学到过知识点分分钟忘得一干二净，今后无论学什么，都做好笔记吧。

## 目录结构

- `assets/images`: 笔记配图
- `assets/templates`: 笔记模板
- `docs`: 基础教程，成体系的，或者分类的
{%- if enable_library_module %}
- `libraries`: 常用库笔记
  - `libraries/standard`: 标准库
  - `libraries/tripartite`: 第三方库
{%- endif %}
- `quickstart`: 基础用法、简介
- `src`: 源码/配置示例
  - `src/docs`: docs 模块示例
  - `src/quickstart`: quickstart 模块示例
{%- if enable_library_module %}
  - `src/libraries/standard`: 标准库示例
  - `src/libraries/tripartite`: 第三方库示例
{%- endif %}
- `tools`: 常用工具笔记
